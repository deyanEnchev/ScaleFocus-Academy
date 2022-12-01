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
	repo    Repository
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
	Id    int    `json:"id"`
}

type Repository interface {
	SaveStories(stories []Story)
}

func NewStoryService(url string, repo Repository) *StoryService {
	return &StoryService{urlBase: url, repo: repo}
}

func (ss *StoryService) FetchTopStories(maxCount int) []int {
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
	arr = arr[:maxCount]

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

	stories := make([]Story, 0)

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

		stories = append(stories, story)
	}
	ss.repo.SaveStories(stories)
	return stories
}
