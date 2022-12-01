package persistence

import (
	"database/sql"
	"final/cmd/persistence/sqlc"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

const (
	LOGIN_INFO = `CREATE TABLE IF NOT EXISTS login_info
	(
		id integer NOT NULL,
		username VARCHAR(256) NOT NULL UNIQUE,
		password VARCHAR(256) NOT NULL,
		PRIMARY KEY (id)
	);`
	USER_LISTS = `CREATE TABLE IF NOT EXISTS user_lists
	(
		id integer NOT NULL,
		user_id integer NOT NULL,
		name VARCHAR(256) NOT NULL,
		PRIMARY KEY (id)
	);`
	USER_TASKS = `CREATE TABLE IF NOT EXISTS user_tasks
	(
		id integer NOT NULL,
		list_id integer NOT NULL,
		text VARCHAR(256) NOT NULL,
		completed bool NOT NULL,
		PRIMARY KEY (id)
	);`
)

func TestDeleteListInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_LISTS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_lists (id, user_id, name) 
	VALUES (?,?,?);`, 1, 1, "John Cena")
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	if assert.NoError(t, repo.DeleteListInDb(1)) {
		rows, _ := mockDb.Query(`SELECT id,user_id,name FROM user_lists WHERE user_id=?;`, 1)
		if rows.Next() {
			t.Fatalf("Delete, not working as intended.")
		}
	}
}

func TestDeleteTaskInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 1, 1, "Workout", false)
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	if assert.NoError(t, repo.DeleteTaskInDb(1)) {
		rows, _ := mockDb.Query(`SELECT * FROM user_tasks WHERE list_id=?;`, 1)
		if rows.Next() {
			t.Fatalf("Delete, not working as intended.")
		}
	}
}

func TestToggleTaskInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 1, 1, "Workout", false)
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	if assert.NoError(t, repo.ToggleTaskInDb(1, true)) {
		rows, _ := mockDb.Query(`SELECT completed FROM user_tasks WHERE id=?;`, 1)
		for rows.Next() {
			var completed bool
			rows.Scan(&completed)
			assert.NotEqual(t, completed, false)
		}
	}
}

func TestCreateListInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_LISTS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	assert.NoError(t, repo.CreateListInDb(1, "Groceries"))
	rows, _ := mockDb.Query(`SELECT * FROM user_lists WHERE user_id=?`, 1)
	assert.NotEmpty(t, rows.Next())
}

func TestCreateTaskInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	assert.NoError(t, repo.CreateTaskInDb(1, "Watermelon", false))
	rows, _ := mockDb.Query(`SELECT * FROM user_tasks WHERE list_id=?;`, 1)
	assert.NotEmpty(t, rows.Next())
}

func TestGetTasksFromDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 1, 1, "Workout", false)
	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 2, 2, "Rest", false)
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))
	got, err := repo.GetTasksFromDb(1)
	want := []sqlc.UserTask{
		{ID: 1, ListID: 1, Text: "Workout", Completed: false},
	}
	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, got, want)
}

func TestGetOneTaskFromDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 1, 1, "Workout", false)
	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 2, 2, "Rest", true)
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))
	got, err := repo.GetOneTaskFromDb(2)
	want := sqlc.UserTask{
		ID: 2, ListID: 2, Text: "Rest", Completed: true,
	}
	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, got, want)
}

func TestGetListsFromDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(USER_LISTS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO user_lists (id, user_id, name) 
	VALUES (?,?,?);`, 1, 1, "John Cena")
	_, err = mockDb.Exec(`INSERT INTO user_lists (id, user_id, name) 
	VALUES (?,?,?);`, 2, 2, "Rey Misterio")
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))
	got, err := repo.GetListsFromDb(1)
	want := []sqlc.UserList{
		{ID: 1, UserID: 1, Name: "John Cena"},
	}
	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, got, want)
}

func TestGetUsersFromDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(LOGIN_INFO)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO login_info (id, username, password) 
	VALUES (?,?,?);`, 1, "John Cena", "test1")
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))
	got, err := repo.GetUsersFromDb()
	want := []sqlc.LoginInfo{
		{ID: 1, Username: "John Cena", Password: "test1"},
	}
	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, got, want)
}

func TestCreateUserInDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(LOGIN_INFO)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))

	// Assertions
	assert.NoError(t, repo.CreateUserInDb("John Cena", "test1"))
	rows, _ := mockDb.Query(`SELECT * FROM login_info;`)
	assert.NotEmpty(t, rows.Next())
}

func TestGetUserTasksFromDb(t *testing.T) {
	// Setup
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed opening in-memory sqlite: %v\n", err)
	}

	_, err = mockDb.Exec(LOGIN_INFO)
	_, err = mockDb.Exec(USER_LISTS)
	_, err = mockDb.Exec(USER_TASKS)
	if err != nil {
		t.Fatalf("Failed creating table: %v\n", err)
	}

	_, err = mockDb.Exec(`INSERT INTO login_info (id, username, password) 
	VALUES (?,?,?);`, 1, "username", "secret1")
	_, err = mockDb.Exec(`INSERT INTO user_lists (id, user_id, name) 
	VALUES (?,?,?);`, 1, 1, "Groceries")
	_, err = mockDb.Exec(`INSERT INTO user_tasks (id, list_id, text, completed)
	VALUES (?,?,?,?);`, 1, 1, "Cucumber", false)
	if err != nil {
		t.Fatalf("Failed inserting into mock database: %v\n", err)
	}

	repo := NewRepository(mockDb, sqlc.New(mockDb))
	got, err := repo.GetUserTasksFromDb(1)
	want := []sqlc.GetUserTasksRow{
		{Name: "Groceries", Text: "Cucumber", Completed: false},
	}

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, got, want)
}
