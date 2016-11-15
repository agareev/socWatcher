package main

import (
	"fmt"
	"os"
	"regexp"
)

// ResponseSocPageList json from 2ch
type ResponseSocPageList struct {
	Threads []struct {
		Files    int    `json:"files_count"`
		TreadNum string `json:"thread_num"`
	} `json:"threads"`
}

// ResponseSocPage json from 2ch
type ResponseSocPage struct {
	Threads []struct {
		Posts []struct {
			Comment string `json:"comment"`
			Num     int    `json:"num"`
		} `json:"posts"`
	} `json:"threads"`
}

var url = "https://2ch.hk/soc/index.json"

func main() {
	regex := os.Args[1]
	streamRegex := regexp.MustCompile(regex)
	dateTread := GetThreadNumber(url)
	threadurl := "https://2ch.hk/soc/res/" + dateTread + ".json"

	page := GetThreadPage(threadurl)
	for number := range page.Threads[0].Posts {
		match := streamRegex.FindStringSubmatch(page.Threads[0].Posts[number].Comment)
		if match != nil {
			fmt.Println(page.Threads[0].Posts[number].Comment)
		}
	}
}
