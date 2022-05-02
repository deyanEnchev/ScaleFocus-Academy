package repository

import (
	"SFA/week9/Lecture26/story"
	"database/sql"
	"log"
	"time"
)

type InputDB struct {
	ID         int
	Title      string
	Score      int
	TimeStored string
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetLastStoryTimeStamp() string {
	query := "SELECT time_stored FROM top_stories"
	var timestamp string
	r.db.QueryRow(query).Scan(&timestamp)
	return timestamp
}

func (r *repository) GetStories() []story.Story {
	var RespDB []story.Story
	query := "SELECT id,title,score FROM top_stories"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var inp story.Story
		if err = rows.Scan(&inp.Id, &inp.Title, &inp.Score); err != nil {
			log.Fatal(err)
		}
		RespDB = append(RespDB, inp)
	}
	return RespDB
}

func (r *repository) SaveStories(stories []story.Story) {
	query := "INSERT INTO top_stories(id,title,score,time_stored) VALUES($1,$2,$3,$4) ON CONFLICT (id) DO UPDATE SET time_stored = $4;"

	for i, s := range stories {
		r.db.Exec(query, i+1, s.Title, s.Score, time.Now().Format(time.RFC1123))
	}
}

// func (r *repository) CheckDB() {
// 	rows, err := r.db.Query(`SELECT * FROM top_stories`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// defer rows.Close()

// 	if !rows.Next() {
// 		r.InitialFillingOfDb()
// 		return
// 	}

// 	rows, _ = r.db.Query(`SELECT * FROM top_stories`)

// 	for rows.Next() {
// 		var inp InputDB
// 		err = rows.Scan(&inp.ID, &inp.Title, &inp.Score, &inp.TimeStored)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		t, err := time.Parse(time.RFC1123, inp.TimeStored)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// timeStored := "Mon, 18 Apr 2022 13:07:31 EEST"
// 		if time.Since(t) > time.Hour {
// 			r.PopulateDb()
// 			return
// 		}
// 		r.RespDB = append(r.RespDB, inp)

// 	}
// }

// func (r *repository) InitialFillingOfDb() {
// 	for i, story := range story.Result.TopStories {
// 		t := time.Now().Format(time.RFC1123)

// 		input := InputDB{i + 1, story.Title, story.Score, t}

// 		r.RespDB = append(r.RespDB, input)
// 		_, err := r.db.Exec(`
// 		INSERT INTO top_stories(id,title,score,time_stored)
// 		VALUES($1,$2,$3,$4); `, input.ID, input.Title, input.Score, input.TimeStored)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func (r *repository) PopulateDb() {
// 	for i, story := range story.Result.TopStories {
// 		t := time.Now().Format(time.RFC1123)

// 		input := InputDB{i + 1, story.Title, story.Score, t}

// 		r.RespDB = append(r.RespDB, input)
// 		_, err := r.db.Exec(`
// 		UPDATE top_stories SET title = $1,
// 		score = $2,
// 		time_stored = $3
// 		WHERE id = $4; `, input.Title, input.Score, input.TimeStored, input.ID)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
