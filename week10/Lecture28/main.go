package main

import (
	"SFA/week10/Lecture28/sqlcDB"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
	Id    int    `json:"id"`
}

type ResponseJSON struct {
	Title      string
	TopStories []Story `json:"top_stories"`
}

type InputDB struct {
	ID         int
	Title      string
	Score      int
	TimeStored string
}

type CreateStoryParams struct {
	ID         int32
	Title      sql.NullString
	Score      sql.NullInt32
	TimeStored sql.NullString
}

var result ResponseJSON
var dbResp []InputDB

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

func HandleJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := json.MarshalIndent(dbResp, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}

const (
	username = "root"
	password = "Asdasdwe123?"
	port     = 3306
	table    = "demo"
)

func CheckDB() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", username, password, port, table))
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	defer db.Close()
	sqlcDB := sqlcDB.New(db)
	ctx := context.Background()
	rows, err := sqlcDB.ListStories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if len(rows) == 0 {
		PopulateDb(sqlcDB)
		return
	}

	for _, v := range rows {
		t, err := time.Parse(time.RFC1123, v.TimeStored.String)
		if err != nil {
			log.Fatal(err)
		}
		if time.Since(t) > time.Minute {
			PopulateDb(sqlcDB)
			return
		}
		inp := InputDB{int(v.ID), v.Title.String, int(v.Score.Int32), t.String()}
		dbResp = append(dbResp, inp)
	}
}

func PopulateDb(sqlcData *sqlcDB.Queries) {
	for _, story := range result.TopStories {
		t := time.Now().Format(time.RFC1123)

		input := InputDB{story.Id, story.Title, story.Score, t}
		dbResp = append(dbResp, input)

		err := sqlcData.CreateStory(context.Background(), sqlcDB.CreateStoryParams{
			ID:           int32(story.Id),
			Title:        sql.NullString{String: story.Title, Valid: true},
			Score:        sql.NullInt32{Int32: int32(story.Score), Valid: true},
			TimeStored:   sql.NullString{String: t, Valid: true},
			TimeStored_2: sql.NullString{String: t, Valid: true},
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	response := Story{}
	response.Worker()
	mux := http.NewServeMux()
	CheckDB()
	mux.Handle("/api/top", HandleJSON())
	http.ListenAndServe(":9000", mux)
}
