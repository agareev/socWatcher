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
			Number  int    `json:"number"`
			Files   []struct {
				Path  string `json:"path"`
				Thumb string `json:"thumbnail"`
			} `json:"files"`
		} `json:"posts"`
	} `json:"threads"`
}

var url = "https://2ch.hk/soc/index.json"
var regex string

func main() {
	if len(os.Args) > 1 {
		fmt.Println(len(os.Args))
		regex = os.Args[1]
	} else {
		regex = ".*"
	}
	streamRegex := regexp.MustCompile(regex)
	dateTread := GetThreadNumber(url)
	threadurl := "https://2ch.hk/soc/res/" + dateTread + ".json"

	page := GetThreadPage(threadurl)
	for number := range page.Threads[0].Posts {
		match := streamRegex.FindStringSubmatch(page.Threads[0].Posts[number].Comment)
		if match != nil {
			push2Database(page.Threads[0].Posts[number].Comment, page.Threads[0].Posts[number].Num)
			// fmt.Println(page.Threads[0].Posts[number].Files[0].Path)
		}
	}
}
