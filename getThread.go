package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetThreadNumber download index page and return number of soc thread
func GetThreadNumber(url string) string {
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
	return page.Threads[0].TreadNum
}

// GetThreadPage - page downloader
func GetThreadPage(url string) *ResponseSocPage {
	threadresp, err := http.Get(url)
	decoder := json.NewDecoder(threadresp.Body)
	page := new(ResponseSocPage)
	ferr := decoder.Decode(&page)
	if ferr != nil {
		fmt.Println("error:", err)
	}
	return page
}
