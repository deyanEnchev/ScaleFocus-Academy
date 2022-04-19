package main

import (
	"SFA/week7/Lecture22/Task2/story"
	"encoding/json"
	"net/http"
)


func HandleJSON(stList []story.Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := json.MarshalIndent(stList, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}

func main() {
	ss := story.NewStoryService("https://hacker-news.firebaseio.com")
	ids := ss.FetchTop10()
	stList := ss.FetchItems(ids)

	mux := http.NewServeMux()
	mux.Handle("/api/top", HandleJSON(stList))
	http.ListenAndServe(":9000", mux)
}
