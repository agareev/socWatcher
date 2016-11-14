package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Response json from 2ch
type ResponseSocPageList struct {
	Threads []struct {
		Files    int    `json:"files_count"`
		TreadNum string `json:"thread_num"`
	} `json:"threads"`
}

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
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(resp.Body)
	page := new(ResponseSocPageList)
	zerr := decoder.Decode(&page)
	if zerr != nil {
		fmt.Println("error:", err)
	}
	dateTread := page.Threads[0].TreadNum

	threadresp, err := http.Get("https://2ch.hk/soc/res/" + dateTread + ".json")
	mecoder := json.NewDecoder(threadresp.Body)
	lage := new(ResponseSocPage)
	ferr := mecoder.Decode(&lage)
	if ferr != nil {
		fmt.Println("error:", err)
	}
	for number := range lage.Threads[0].Posts {
		fmt.Println(lage.Threads[0].Posts[number].Comment)
	}
}
