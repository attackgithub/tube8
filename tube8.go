package tube8

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "http://api.tube8.com/api.php"
const APITimeout = 3

func SearchVideos(search string) Tube8SearchResult {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"?action=searchVideos&output=xml&search=%s&thumbsize=all", url.QueryEscape(search)))
	b, _ := ioutil.ReadAll(resp.Body)
	var result Tube8SearchResult
	err := xml.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetVideoByID(ID string) Tube8SingleVideo {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"?action=getvideobyid&video_id=%s&output=xml&thumbsize=all", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result Tube8SingleVideo
	err := xml.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetVideoEmbedCode(ID string) Tube8EmbedCode {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"?action=getvideoembedcode&output=xml&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result Tube8EmbedCode
	err := xml.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
