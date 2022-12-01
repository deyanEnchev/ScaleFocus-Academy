package main

import (
	"SFA/week9/Lecture26/handlers"
	"SFA/week9/Lecture26/repository"
	"SFA/week9/Lecture26/story"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdasdwe123"
	dbname   = "demo"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	defer db.Close()
	mux := http.NewServeMux()
	repo := repository.NewRepository(db)
	api := handlers.API{
		StorageService: repo,
		StoryService:   *story.NewStoryService("https://hacker-news.firebaseio.com", repo),
	}
	mux.Handle("/api/top", api.HandleTopStories())
	http.ListenAndServe(":9000", mux)
}
