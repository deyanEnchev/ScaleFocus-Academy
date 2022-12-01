package repository

import (
	"SFA/week9/Lecture26/story"
	"database/sql"
	"reflect"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

const (
	createTable = "CREATE TABLE IF NOT EXISTS top_stories (id INT PRIMARY KEY, title TEXT, score INT, time_stored TEXT)"
	insert      = "INSERT INTO top_stories (id,title,score,time_stored) VALUES(?,?,?,?)"
)

func TestGetLastStoryTimeStamp(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in-memory DB")
	}

	_, err = mockDb.Exec(createTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)
	result := repo.GetLastStoryTimeStamp()
	if result == time.Now().Format(time.RFC1123) {
		t.Fatal("Failed to create initial condition")
	}

	wantedTime := time.Now().Add(time.Hour).Format(time.RFC1123)
	mockDb.Exec(insert, 0, "UnitTest", 15, wantedTime)
	mockDb.Exec(insert, 1, "UnitTest1", 15, time.Now().Add(-time.Hour).Format(time.RFC1123))

	result = repo.GetLastStoryTimeStamp()
	if result != wantedTime {
		t.Fatal("Failed to get latest timestamp")
	}
}

func TestGetStories(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in-memory DB")
	}

	_, err = mockDb.Exec(createTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	mockDb.Exec(insert, 1, "UnitTest1", 15, time.Now().Add(time.Minute).Format(time.RFC1123))
	mockDb.Exec(insert, 2, "UnitTest2", 25, time.Now().Add(-time.Hour).Format(time.RFC1123))
	

	repo := NewRepository(mockDb)
	got := repo.GetStories()
	want := []story.Story{
		{Title: "UnitTest1",Score: 15, Id: 1},
		{Title: "UnitTest2",Score: 25, Id: 2},
	}
	
	if !reflect.DeepEqual(got,want) {
		t.Fatalf("Test failed, wanted: %v, got: %v", want,got)
	}
}

func TestSaveStories(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in-memory DB")
	}

	_, err = mockDb.Exec(createTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	stList := []story.Story{
		{Title: "UnitTest1",Score: 15, Id: 1},
		{Title: "UnitTest2",Score: 25, Id: 2},
	}
	repo := NewRepository(mockDb)
	repo.SaveStories(stList)
	got := repo.GetStories()
	
	if !reflect.DeepEqual(got,stList) {
		t.Fatalf("Test failed, wanted: %v, got: %v", stList, got)
	}

}
