package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "snapchat-dl",
	Short: "Snapchat-dl is a tool to download Snapchat videos",
	Long:  "Snapchat-dl is a tool to download Snapchat videos",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var userName = args[0]
		fmt.Println("Fetching data for " + userName)

		_, stories, _, _ := fetchStories(userName)

		if len(stories) == 0 {
			fmt.Printf("%s has no stories\n", userName)
			return
		}

		for _, story := range stories {
			snapId := story.SnapID.Value
			mediaUrl := story.SnapUrls.MediaURL
			mediaType := story.SnapMediaType
			timestamp, _ := strconv.Atoi(story.TimestampInSec.Value)
			dateStr := time.Unix(int64(timestamp), 0).Format("2006-01-02")

			filename := fmt.Sprintf("%s_%s.%s", snapId, userName, MEDIA_TYPE[mediaType])

			directory, err := os.Getwd()
			checkIsError(err)

			filePath := fmt.Sprintf("%s/%s/%s/%s", directory, userName, dateStr, filename)

			go downloadMedia(mediaUrl, filePath, 0)
		}
	},
}

func downloadMedia(url string, destination string, interval int) {
	dir := filepath.Dir(destination)
	if len(dir) > 0 {
		os.MkdirAll(dir, 0755)
	}

	time.Sleep(time.Duration(interval))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
	checkIsError(err)
	defer resp.Body.Close()

	fmt.Printf(resp.Status)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download: status code %d\n", resp.StatusCode)
		return
	}

	file, exists := alreadyExists(destination)

	if exists && file.Size() == resp.ContentLength {
		fmt.Println("file already exists")
		return
	}

	if exists && file.Size() == 0 {
		err := os.Remove(destination)
		checkIsError(err)
	}

	// Create the destination file
	out, err := os.Create(destination)
	checkIsError(err)
	defer out.Close()

	// Copy the response body directly to the file
	_, err = io.Copy(out, resp.Body)
	checkIsError(err)

	fmt.Println("Downloaded " + destination)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Snapchat-dl",
	Long:  "All software has versions. This is Snapchat-dl's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Snapchat-dl v" + version)
	},
}
