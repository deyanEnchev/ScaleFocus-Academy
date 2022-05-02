package story

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

type FakeStorage struct {
	savedStories []Story
}

func (fs *FakeStorage) SaveStories(stories []Story) {
	fs.savedStories = stories
}

func TestHandleTopStories(t *testing.T) {
	want := []int{5283958, 4283958, 3283958, 2283958, 1283958, 283958, 183958, 83958, 73958, 63958}

	router := http.NewServeMux()
	mockServer := httptest.NewServer(router)

	router.Handle("/v0/topstories.json", MockTopStoriesHandler(want))
	worker := NewStoryService(mockServer.URL, &FakeStorage{})

	got := worker.FetchTopStories(len(want))

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}

}

func MockTopStoriesHandler(idList []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(idList)
	}
}

func TestHandleItem(t *testing.T) {
	s1 := Story{Title: "Test1", Score: 51, Id: 5283958}
	s2 := Story{Title: "Test2", Score: 512, Id: 4283958}
	s3 := Story{Title: "Test3", Score: 5123, Id: 3283958}
	s4 := Story{Title: "Test4", Score: 51234, Id: 2283958}
	s5 := Story{Title: "Test5", Score: 512345, Id: 1283958}
	s6 := Story{Title: "Test6", Score: 5123456, Id: 283958}
	s7 := Story{Title: "Test7", Score: 5654321, Id: 183958}
	s8 := Story{Title: "Test8", Score: 56543, Id: 83958}
	s9 := Story{Title: "Test9", Score: 5654, Id: 73958}
	s10 := Story{Title: "Test10", Score: 565, Id: 63958}

	stories := []Story{s1, s2, s3, s4, s5, s6, s7, s8, s9, s10}

	router := http.NewServeMux()
	mockServer := httptest.NewServer(router)
	router.Handle("/v0/item/", MockItemHandler(stories, mockServer.URL))
	fs := &FakeStorage{}
	worker := NewStoryService(mockServer.URL, fs)
	defer mockServer.Close()
	idList := []int{5283958, 4283958, 3283958, 2283958, 1283958, 283958, 183958, 83958, 73958, 63958}
	got2 := worker.FetchItems(idList)

	if !reflect.DeepEqual(got2, stories) {
		t.Fatalf(`Got %v, want %v.`, got2, stories)
	}

	if !reflect.DeepEqual(got2, fs.savedStories) {
		t.Fatalf(`Got %v, want %v.`, got2, fs.savedStories)
	}

}

func MockItemHandler(stories []Story, mockUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(mockUrl)

		if err != nil {
			log.Fatal(err)
		}
		for _, story := range stories {
			str := strconv.Itoa(story.Id)
			u.Path = "/v0/item/" + str + ".json"

			w.WriteHeader(http.StatusOK)
			body, err := json.Marshal(story)
			if err != nil {
				log.Fatal(err)
			}
			stories = stories[1:]
			w.Write(body)
			return

		}
	}
}
