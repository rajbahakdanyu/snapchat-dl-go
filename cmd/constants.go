package cmd

const (
	version       = "0.1.0"
	baseUrl       = "https://www.snapchat.com/add/"
	regexpWebJson = `<script\s*id="__NEXT_DATA__"\s*type="application\/json">([^<]+)<\/script>`
)

var mediaTypes = []string{"jpg", "mp4"}

var quiet bool
var maxStoryNum uint16
var sleepInterval uint16
var maxWorkers uint16

const helpTemplate = `usage: go run . [-h] [-l MAX_NUM_STORY] [-j MAX_WORKERS] [-t INTERVAL] [--n INTERVAL]
                   [-q]
                   [username ...]

positional arguments:
  username              At least one or more usernames to download stories for.

options:
  -h, --help            show this help message and exit
  -l MAX_NUM_STORY, --limit-story MAX_NUM_STORY
                        Set maximum number of stories to download.
  -j MAX_WORKERS, --max-concurrent-downloads MAX_WORKERS
                        Set maximum number of parallel downloads.
  -t INTERVAL, --update-interval INTERVAL
                        Set the update interval for checking new story in seconds. (Default: 10m)
  --n --sleep-interval INTERVAL
                        Sleep between downloads in seconds. (Default: 1s)
  -q, --quiet           Do not print anything except errors to the console.
`
