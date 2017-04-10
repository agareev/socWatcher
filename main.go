package main

import (
	"encoding/json"
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
			Num     int    `json:"num"`    // номер поста в этом треде (относительный)
			Number  int    `json:"number"` // номер поста (абсолютный)
			Files   []struct {
				Path  string `json:"path"`
				Thumb string `json:"thumbnail"`
			} `json:"files"`
		} `json:"posts"`
	} `json:"threads"`
}

// Config is type to parse config.json
type Config struct {
	DBIP    string
	DBLogin string
	DBPass  string
}

var url = "https://2ch.hk/soc/index.json"
var regex string
var configuration Config

// ReadConfig is function for read jsonconf
func ReadConfig(path string) Config {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	configuration := new(Config)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return *configuration
}

func getRegex() string {
	if len(os.Args) > 1 {
		fmt.Println(os.Args)
		return os.Args[1]
	}
	return ".*"
}

func outputComments(page *ResponseSocPage, streamRegex *regexp.Regexp) map[int]string {
	comments := make(map[int]string)
	for number := range page.Threads[0].Posts {
		match := streamRegex.FindStringSubmatch(page.Threads[0].Posts[number].Comment)
		if match != nil {
			// push2Database(db, page.Threads[0].Posts[number].Comment, page.Threads[0].Posts[number].Num)
			id := page.Threads[0].Posts[number].Num
			comment := page.Threads[0].Posts[number].Comment
			comments[id] = comment
		}
	}
	return comments
}

func init() {
	configuration = ReadConfig("config.json")
}

func main() {
	regex := getRegex()
	streamRegex := regexp.MustCompile(regex)

	dateTread := GetThreadNumber(url)
	threadurl := "https://2ch.hk/soc/res/" + dateTread + ".json"
	page := GetThreadPage(threadurl)
	push2Database(outputComments(page, streamRegex))
}
