package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

type ResponseJSON struct {
	TopStories []Story `json:"top_stories"`
}

var result ResponseJSON

func (rj *Story) Worker() {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	var arr []int
	err = json.Unmarshal(body, &arr)

	if err != nil {
		log.Fatal(err)
	}
	arr = arr[:10]

	slice := make([]string, 0, 10)
	for _, v := range arr {
		slice = append(slice, strconv.Itoa(v))
	}

	u, err := url.Parse("https://hacker-news.firebaseio.com")
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
}

func HandleUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := json.MarshalIndent(result, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}

func main() {
	response := Story{}
	response.Worker()
	mux := http.NewServeMux()
	mux.Handle("/top", HandleUser())
	http.ListenAndServe(":9000", mux)
}
