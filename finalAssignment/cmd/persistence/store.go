package persistence

import (
	"context"
	"database/sql"
	"final/cmd/persistence/sqlc"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type repository struct {
	db *sql.DB
	*sqlc.Queries
}

func NewRepository(db *sql.DB, s *sqlc.Queries) *repository {
	return &repository{db, s}
}

func ConnectToDb() (*sql.DB, *sqlc.Queries) {
	username := os.Getenv("USERNAME_DB")
	password := os.Getenv("PASSWORD_DB")
	port := os.Getenv("PORT_DB")
	table := os.Getenv("TABLE_DB")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", username, password, port, table))
	if err != nil {
		fmt.Printf("Failed connecting to the database: %v", err)
		return nil, nil
	}
	s := sqlc.New(db)
	return db, s
}

func (r *repository) DeleteListInDb(id int) error {
	err := r.DeleteList(context.Background(), int32(id))
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTaskInDb(id int) error {
	err := r.DeleteTask(context.Background(), int32(id))
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ToggleTaskInDb(id int, compl bool) error {
	err := r.ToggleTask(context.Background(), sqlc.ToggleTaskParams{
		Completed: compl,
		ID:        int32(id),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateListInDb(userId int, name string) error {
	err := r.CreateList(context.Background(), sqlc.CreateListParams{
		UserID: int32(userId),
		Name:   name,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateTaskInDb(listId int, text string, completed bool) error {
	err := r.CreateTask(context.Background(), sqlc.CreateTaskParams{
		ListID:    int32(listId),
		Text:      text,
		Completed: completed,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTasksFromDb(listId int) ([]sqlc.UserTask, error) {
	t, err := r.GetTasks(context.Background(), int32(listId))
	if err != nil {
		return t, err
	}
	return t, nil
}

func (r *repository) GetOneTaskFromDb(taskId int) (sqlc.UserTask, error) {
	t, err := r.GetTask(context.Background(), int32(taskId))
	if err != nil {
		return t, err
	}
	return t, nil
}

func (r *repository) GetListsFromDb(userId int) ([]sqlc.UserList, error) {
	l, err := r.GetLists(context.Background(), int32(userId))
	if err != nil {
		return l, err
	}
	return l, nil
}

func (r *repository) GetUsersFromDb() ([]sqlc.LoginInfo, error) {
	res, err := r.GetUsers(context.Background())
	if err != nil {
		return []sqlc.LoginInfo{}, err
	}

	return res, nil
}

func (r *repository) CreateUserInDb(username, password string) error {
	err := r.SaveLoginInfo(context.Background(), sqlc.SaveLoginInfoParams{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserTasksFromDb(userId int) ([]sqlc.GetUserTasksRow, error) {
	t, err := r.GetUserTasks(context.Background(), int32(userId))
	if err != nil {
		return t, err
	}
	return t, nil
}
