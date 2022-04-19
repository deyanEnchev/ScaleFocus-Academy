package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdasdwe123"
	dbname   = "demo"
)

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

func CheckDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM top_stories`)
	if err != nil {
		log.Fatal(err)
	}

	// defer rows.Close()
	if !rows.Next() {
		PopulateDb(db)
		return
	}

	for rows.Next() {
		var inp InputDB
		err = rows.Scan(&inp.ID, &inp.Title, &inp.Score, &inp.TimeStored)

		if err != nil {
			log.Fatal(err)
		}

		t, err := time.Parse(time.RFC1123, inp.TimeStored)

		if err != nil {
			log.Fatal(err)
		}
		// timeStored := "Mon, 18 Apr 2022 13:07:31 EEST"
		if time.Since(t) > time.Hour {
			PopulateDb(db)
			return
		}
		dbResp = append(dbResp, inp)

	}
}

func PopulateDb(db *sql.DB) {
	for i, story := range result.TopStories {
		t := time.Now().Format(time.RFC1123)

		input := InputDB{i, story.Title, story.Score, t}

		dbResp = append(dbResp, input)
		_, err := db.Exec(`
		UPDATE top_stories SET id = $1,
		title = $2,
		score = $3,
		time_stored = $4; `, input.ID, input.Title, input.Score, input.TimeStored)

		if err != nil {
			log.Fatal(err)
		}
	}
}

//INSERT INTO top_stories(id,title,score,time_stored)
//VALUES($1,$2,$3,$4);

func main() {
	response := Story{}
	response.Worker()
	mux := http.NewServeMux()
	CheckDB()
	mux.Handle("/api/top", HandleJSON())
	http.ListenAndServe(":9000", mux)
}
