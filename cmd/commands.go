package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
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

		var wg sync.WaitGroup
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

			wg.Add(1)
			go func(url, path string) {
				defer wg.Done()
				downloadMedia(url, path, 0)
			}(mediaUrl, filePath)
		}
		wg.Wait()
	},
}

func downloadMedia(url string, destination string, interval int) {
	dir := filepath.Dir(destination)
	if len(dir) > 0 {
		err := os.MkdirAll(dir, 0755)
		checkIsError(err)
	}

	time.Sleep(time.Duration(interval) * time.Second)

	fmt.Printf("Requesting %s\n", url)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download: status code %d\n", resp.StatusCode)
		return
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Content-Length: %d bytes\n", resp.ContentLength)

	// Check if file already exists
	if fileInfo, err := os.Stat(destination); err == nil {
		if fileInfo.Size() == resp.ContentLength {
			fmt.Printf("File already exists and has the correct size: %s\n", destination)
			return
		}
	}

	// Create the destination file
	out, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	// Copy the response body directly to the file
	bytesWritten, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Downloaded %s (%d bytes written)\n", destination, bytesWritten)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Snapchat-dl",
	Long:  "All software has versions. This is Snapchat-dl's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Snapchat-dl v" + version)
	},
}
