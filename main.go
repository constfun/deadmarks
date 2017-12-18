package main

import (
    "encoding/json"
    "log"
	"io/ioutil"
	"os"
	"fmt"
	"strings"
	"net/http"
)

type Bookmark struct {
	Type string
    Title string
    Uri *string
    Children []Bookmark
}

func (b Bookmark) CollectUris(uris *map[string]bool) {
	if b.Uri != nil &&
		(strings.HasPrefix(*b.Uri, "http:") ||
			strings.HasPrefix(*b.Uri,"https:")) {

		(*uris)[*b.Uri] = true
	}
	if b.Children != nil {
		for _, child := range b.Children {
			child.CollectUris(uris)
		}
	}
}

func TestUri(uri string) bool {
	resp, err := http.Get(uri)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Printf("File read err: %v\n", err)
		return
	}

	var bmark Bookmark
	err = json.Unmarshal(file, &bmark)
	if err != nil {
		log.Printf("Unmarshal err: %v", err)
		return
	}

	uris := make(map[string]bool)
	bmark.CollectUris(&uris)

	for uri := range uris {
		res := TestUri(uri)
		if res {
			fmt.Printf("OK: %v\n", uri)
		} else {
			fmt.Printf("DOWN: %v\n", uri)
		}
	}
}