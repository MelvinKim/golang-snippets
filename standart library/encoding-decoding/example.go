package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document from the search
	gResult struct {
		GSearchResultClass string `json:"GSearchResultClass"`
		UnescapedURL       string `json:"UnescapedURL"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNotFormatting string `json:"titleNotFormatting"`
		Content            string `json:"content"`
	}

	// gResponse contains the top level document
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	defer resp.Body.Close()

	// decode the JSON response to our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("gResponse: ", gr)
}
