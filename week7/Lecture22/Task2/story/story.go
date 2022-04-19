package story

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type StoryService struct {
	urlBase string
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
	Id    int    `json:"id"`
}

type ResponseJSON struct {
	TopStories []Story `json:"top_stories"`
}

var result ResponseJSON

func NewStoryService(url string) *StoryService {
	return &StoryService{urlBase: url}
}

func (ss *StoryService) FetchTop10() []int {
	res, err := http.Get(ss.urlBase + "/v0/topstories.json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var arr []int
	err = json.Unmarshal(body, &arr)

	if err != nil {
		log.Fatal(err)
	}
	arr = arr[:10]

	return arr
}

func (ss *StoryService) FetchItems(top10 []int) []Story {
	slice := make([]string, 0, 10)
	for _, v := range top10 {
		slice = append(slice, strconv.Itoa(v))
	}

	u, err := url.Parse(ss.urlBase)
	if err != nil {
		log.Fatal(err)
	}

	for _, detail := range slice {
		u.Path = "/v0/item/" + detail + ".json"

		response, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var story Story

		err = json.Unmarshal(body, &story)
		if err != nil {
			log.Fatal(err)
		}

		result.TopStories = append(result.TopStories, story)
	}
	return result.TopStories
}
