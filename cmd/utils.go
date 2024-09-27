package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"regexp"
)

const baseUrl = "https://www.snapchat.com/add/"
const regexpWebJson = `<script\s*id="__NEXT_DATA__"\s*type="application\/json">([^<]+)<\/script>`

var MEDIA_TYPE = []string{"jpg", "mp4"}

func checkIsError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

func alreadyExists(filePath string) (fs.FileInfo, bool) {
	file, err := os.Stat(filePath)
	exists := errors.Is(err, os.ErrExist)
	return file, exists
}

func fetchStories(userName string) (UserProfile, []SnapList, []CuratedHighlights, []SpotlightHighlights) {
	resp, err := http.Get(baseUrl + userName)
	checkIsError(err)

	body, err := io.ReadAll(resp.Body)
	checkIsError(err)

	expression, err := regexp.Compile(regexpWebJson)
	checkIsError(err)

	matches := expression.FindAll(body, -1)
	if len(matches) == 0 {
		fmt.Println("No matches found")
	}

	var formattedBody []byte
	formattedBody = bytes.TrimLeft(matches[0], `<script\s id="__NEXT_DATA__" type="application\/json">`)
	formattedBody = bytes.TrimRight(formattedBody, "</script>")

	var parsedJson SnapchatData
	error := json.Unmarshal(formattedBody, &parsedJson)
	checkIsError(error)

	var userInfo = parsedJson.Props.PageProps.UserProfile
	var stories = parsedJson.Props.PageProps.Story.SnapList
	var curatedHighlights = parsedJson.Props.PageProps.CuratedHighlights
	var spotHighlights = parsedJson.Props.PageProps.SpotlightHighlights

	return userInfo, stories, curatedHighlights, spotHighlights
}
