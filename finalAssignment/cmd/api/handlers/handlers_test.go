package handlers

import (
	"final/cmd/persistence/sqlc"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (m API) DeleteTaskInDb(id int) error {
	return nil
}
func TestDeleteTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	m := API{}

	// Assertions
	if assert.NoError(t, m.DeleteTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func (m API) DeleteListInDb(id int) error {
	return nil
}
func TestDeleteList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	m := API{}

	// Assertions
	if assert.NoError(t, m.DeleteList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func (m API) ToggleTaskInDb(id int, compl bool) error {
	return nil
}

func (m API) GetOneTaskFromDb(taskId int) (sqlc.UserTask, error) {
	result := sqlc.UserTask{
		Completed: true,
	}
	return result, nil
}
func TestToggleTask(t *testing.T) {
	// Setup
	e := echo.New()
	body := `{"text":"","completed":true}
`
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	m := API{}

	// Assertions
	if assert.NoError(t, m.ToggleTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, body, rec.Body.String())
	}
}

func (m API) CreateListInDb(userId int, name string) error {
	return nil
}

func (m API) GetListsFromDb(userId int) ([]sqlc.UserList, error) {
	result := []sqlc.UserList{
		{ID: 1, Name: "test"},
	}
	return result, nil
}
func TestAddList(t *testing.T) {
	// Setup
	e := echo.New()
	body := `{"name":"test"}
`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists")
	c.Set("UserID", 1)
	m := API{}
	expected := `{"id":1,"name":"test"}
`
	// Assertions
	if assert.NoError(t, m.AddList(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func (m API) CreateTaskInDb(listId int, text string, completed bool) error {
	return nil
}

func (m API) GetTasksFromDb(listId int) ([]sqlc.UserTask, error) {
	result := []sqlc.UserTask{
		{Text: "test", ListID: 1},
	}
	return result, nil
}
func TestAddTask(t *testing.T) {
	// Setup
	e := echo.New()
	body := `{"text":"test","listId":1,"completed":false}
`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists/:id/tasks")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("UserID", 1)
	m := API{}

	// Assertions
	if assert.NoError(t, m.AddTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, body, rec.Body.String())
	}
}

func TestGetTasks(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("lists/:id/tasks")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("UserID", 1)
	m := API{}
	expected := `[{"text":"test","listId":1,"completed":false}]
`
	// Assertions
	if assert.NoError(t, m.GetTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetLists(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("lists")
	c.Set("UserID", 1)
	m := API{}
	expected := `[{"id":1,"name":"test"}]
`
	// Assertions
	if assert.NoError(t, m.GetLists(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func (m API) GetUserTasksFromDb(userId int) ([]sqlc.GetUserTasksRow, error) {
	result := []sqlc.GetUserTasksRow{
		{Name: "Groceries", Text: "Test1", Completed: false},
	}
	return result, nil
}
func TestExportLists(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/export")
	c.Set("UserID", 1)
	m := API{}
	expected := "List,Task,Status\nGroceries,Test1,false\n"

	// Assertions
	if assert.NoError(t, m.ExportLists(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetWeather(t *testing.T) {
	// Setup
	e := echo.New()
	os.Setenv("API_KEY", "19f20f0a68aa2000ad6460e17d9f88d4")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add("lat", "42.698334")
	req.Header.Add("lon", "23.319941")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/weather")
	c.Set("UserID", 1)
	c.Set("Weather-API", "https://api.openweathermap.org/data/2.5/weather")

	// Assertions
	if assert.NoError(t, GetWeather(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Sofia")
	}
	os.Unsetenv("API_KEY")
}
