package handlers

import (
	"SFA/week9/Lecture26/story"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type API struct {
	StorageService Storage
	StoryService   story.StoryService
}

type Storage interface {
	GetLastStoryTimeStamp() string
	GetStories() []story.Story
}

// respDB []repository.InputDB
func (api API) HandleTopStories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		t := time.Now().Add(-time.Hour)
		if !(api.StorageService.GetLastStoryTimeStamp() == "") {
			var err error
			t, err = time.Parse(time.RFC1123, api.StorageService.GetLastStoryTimeStamp())
			if err != nil {
				log.Fatal(err)
			}
		}

		var stList []story.Story
		if time.Since(t) > time.Hour {
			temp := api.StoryService.FetchTopStories(10)
			stList = api.StoryService.FetchItems(temp)
		} else {
			stList = api.StorageService.GetStories()
		}

		resp, err := json.MarshalIndent(stList, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}
