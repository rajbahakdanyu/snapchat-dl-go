package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func fileExists(path string, expectedSize int64) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.Size() == expectedSize
}

func fetchStories(userName string) (UserProfile, []SnapList, []CuratedHighlights, []SpotlightHighlights) {
	resp, err := http.Get(baseUrl + userName)
	checkError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError(err)

	re := regexp.MustCompile(regexpWebJson)
	matches := re.FindSubmatch(body)
	if len(matches) < 2 {
		fmt.Println("No matches found")
		return UserProfile{}, nil, nil, nil
	}

	var parsedJson SnapchatData
	err = json.Unmarshal(matches[1], &parsedJson)
	checkError(err)

	return parsedJson.Props.PageProps.UserProfile,
		parsedJson.Props.PageProps.Story.SnapList,
		parsedJson.Props.PageProps.CuratedHighlights,
		parsedJson.Props.PageProps.SpotlightHighlights
}

func runRoot(cmd *cobra.Command, args []string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	for _, userName := range args {
		if !quiet {
			fmt.Printf("Fetching data for %s\n", userName)
			s.Start()
		}

		_, stories, _, _ := fetchStories(userName)

		if len(stories) == 0 {
			s.FinalMSG = fmt.Sprintf("%s has no stories\n", userName)
			s.Stop()
			continue
		}

		numOfStories := int(maxStoryNum)
		if maxStoryNum == 0 || numOfStories > len(stories) {
			numOfStories = len(stories)
		}

		var wg sync.WaitGroup
		for i := 0; i < numOfStories; i++ {
			wg.Add(1)
			go func(s SnapList) {
				defer wg.Done()
				downloadStory(s, userName)
			}(stories[i])

			time.Sleep(time.Duration(sleepInterval) * time.Second)
		}
		wg.Wait()

		s.FinalMSG = fmt.Sprintf("Downloaded %d stories for %s\n", numOfStories, userName)
		s.Stop()
	}
}

func downloadStory(story SnapList, userName string) {
	snapId := story.SnapID.Value
	mediaUrl := story.SnapUrls.MediaURL
	mediaType := story.SnapMediaType
	timestamp, _ := strconv.ParseInt(story.TimestampInSec.Value, 10, 64)
	dateStr := time.Unix(timestamp, 0).Format("2006-01-02")

	filename := fmt.Sprintf("%s_%s.%s", snapId, userName, mediaTypes[mediaType])
	directory, err := os.Getwd()
	checkError(err)

	filePath := filepath.Join(directory, userName, dateStr, filename)
	downloadMedia(mediaUrl, filePath)
}

func downloadMedia(url, destination string) {
	dir := filepath.Dir(destination)
	if dir != "" {
		err := os.MkdirAll(dir, 0755)
		checkError(err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download: status code %d\n", resp.StatusCode)
		return
	}

	if fileExists(destination, resp.ContentLength) {
		fmt.Printf("File already exists: %s\n", destination)
		return
	}

	out, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}
