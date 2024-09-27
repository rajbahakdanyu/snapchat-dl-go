package cmd

const (
	version       = "0.1.0"
	baseUrl       = "https://www.snapchat.com/add/"
	regexpWebJson = `<script\s*id="__NEXT_DATA__"\s*type="application\/json">([^<]+)<\/script>`
)

var mediaTypes = []string{"jpg", "mp4"}

var quiet bool
