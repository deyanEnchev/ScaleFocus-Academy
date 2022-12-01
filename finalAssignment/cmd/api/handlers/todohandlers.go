package handlers

import (
	"encoding/csv"
	"encoding/json"
	"final/cmd/persistence/sqlc"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type API struct {
	Storage
}

type Storage interface {
	DeleteTaskInDb(id int) error
	DeleteListInDb(id int) error
	ToggleTaskInDb(id int, compl bool) error
	CreateListInDb(userId int, name string) error
	CreateTaskInDb(listId int, text string, completed bool) error
	GetListsFromDb(userId int) ([]sqlc.UserList, error)
	GetTasksFromDb(listId int) ([]sqlc.UserTask, error)
	GetOneTaskFromDb(taskId int) (sqlc.UserTask, error)
	GetUserTasksFromDb(userId int) ([]sqlc.GetUserTasksRow, error)
}

type Task struct {
	Id        int    `json:"id,omitempty"`
	Text      string `json:"text"`
	ListId    int    `json:"listId,omitempty"`
	Completed bool   `json:"completed"`
}

type List struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type WeatherInfo struct {
	FormatedTemp string `json:"formatedTemp"`
	Description  string `json:"description"`
	City         string `json:"city"`
}

type weatherData struct {
	City string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type ExportData struct {
	Name      string
	Text      string
	Completed string
}

func GetWeather(c echo.Context) error {
	url, _ := url.Parse(c.Get("Weather-API").(string))
	q := url.Query()
	q.Set("lat", c.Request().Header.Get("lat"))
	q.Set("lon", c.Request().Header.Get("lon"))
	q.Set("appid", os.Getenv("API_KEY"))
	q.Set("units", "metric")
	url.RawQuery = q.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("Failed getting from api endpoint: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)

	w := new(weatherData)
	err = json.Unmarshal(b, &w)
	if err != nil {
		fmt.Printf("Failed unmarshaling weather data: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	temp := strconv.FormatFloat(w.Main.Celsius, 'f', 1, 32) + "°C"
	wi := WeatherInfo{
		FormatedTemp: temp,
		Description:  w.Weather[0].Description,
		City:         w.City,
	}

	return c.JSON(http.StatusOK, wi)
}

func (a API) ExportLists(c echo.Context) error {
	c.Response().Header().Set("content-type", "text/csv")
	t, err := a.GetUserTasksFromDb(c.Get("UserID").(int))
	if err != nil {
		fmt.Printf("Failed fetching user's tasks from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}
	csvFile, err := os.Create("user_tasks.csv")
	if err != nil {
		fmt.Printf("Couldn't create csv file: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	csvWriter := csv.NewWriter(csvFile)

	err = csvWriter.Write([]string{"List", "Task", "Status"})
	if err != nil {
		fmt.Printf("Something went wrong with the writing to the csv file: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	for _, row := range t {
		cb := strconv.FormatBool(row.Completed)
		err = csvWriter.Write([]string{row.Name, row.Text, cb})
		if err != nil {
			fmt.Printf("Something went wrong with the writing to the csv file: %v\n", err)
			return c.String(http.StatusInternalServerError, "")
		}
	}

	csvWriter.Flush()
	csvFile.Close()
	return c.Attachment("user_tasks.csv", "user_tasks.csv")
}

func (a API) DeleteList(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id in deleteList: %v\n", err)
		return c.String(http.StatusInternalServerError, "Invalid id input.")
	}

	err = a.DeleteListInDb(id)
	if err != nil {
		fmt.Printf("Failed deleting list from database: %v\n", err)
		return c.String(http.StatusInternalServerError, "Failed to delete list from database.")
	}

	return c.String(http.StatusOK, "Test")
}

func (a API) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id in deleteTask: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = a.DeleteTaskInDb(id)
	if err != nil {
		fmt.Printf("Failed deleting task from database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, "")
}

func (a API) ToggleTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Printf("Failed reading the request body for addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}
	defer c.Request().Body.Close()

	t := new(Task)
	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Printf("Failed unmarshaling in addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = a.ToggleTaskInDb(id, t.Completed)
	if err != nil {
		fmt.Printf("Failed toggling task in database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	tr, err := a.GetOneTaskFromDb(id)
	if err != nil {
		fmt.Printf("Couldn't get tasks from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	result := Task{
		Id:        int(tr.ID),
		Text:      tr.Text,
		ListId:    int(tr.ListID),
		Completed: tr.Completed,
	}
	return c.JSON(http.StatusOK, result)
}

func (a API) AddList(c echo.Context) error {
	l := new(List)
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Printf("Failed reading the request body for addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Printf("Failed unmarshaling in addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if l.Id != 0 {
		fmt.Printf("Filling in ID is prohibited.\n")
		return c.String(http.StatusInternalServerError, "")
	}

	userId := c.Get("UserID").(int)
	err = a.CreateListInDb(userId, l.Name)
	if err != nil {
		fmt.Printf("Couldn't insert list in database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	list, err := a.GetListsFromDb(userId)
	if err != nil {
		fmt.Printf("Couldn't get lists from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	result := List{
		Id:   int(list[len(list)-1].ID),
		Name: list[len(list)-1].Name,
	}

	return c.JSON(http.StatusCreated, result)
}

func (a API) AddTask(c echo.Context) error {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	t := new(Task)
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Printf("Failed reading the request body for addTask: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Printf("Failed unmarshaling in addTask: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if t.Id != 0 || t.ListId != 0 {
		fmt.Println("Filling in IDs is prohibited.")
	}

	err = a.CreateTaskInDb(listId, t.Text, t.Completed)
	if err != nil {
		fmt.Printf("Couldn't insert task in database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	tasks, err := a.GetTasksFromDb(listId)
	if err != nil {
		fmt.Printf("Couldn't get tasks from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	result := Task{
		Id:        int(tasks[len(tasks)-1].ID),
		Text:      tasks[len(tasks)-1].Text,
		ListId:    int(tasks[len(tasks)-1].ListID),
		Completed: tasks[len(tasks)-1].Completed,
	}

	return c.JSON(http.StatusCreated, result)
}

func (a API) GetTasks(c echo.Context) error {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	result := []Task{}
	res, err := a.GetTasksFromDb(listId)
	if err != nil {
		fmt.Printf("Couldn't get tasks from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	for _, row := range res {
		result = append(result, Task{
			Id:        int(row.ID),
			Text:      row.Text,
			ListId:    int(row.ListID),
			Completed: row.Completed,
		})

	}

	return c.JSON(http.StatusOK, result)
}

func (a API) GetLists(c echo.Context) error {
	l, err := a.GetListsFromDb(c.Get("UserID").(int))
	if err != nil {
		fmt.Printf("Couldn't get lists from the database: %v\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	result := []List{}

	for _, row := range l {
		result = append(result, List{
			Id:   int(row.ID),
			Name: row.Name,
		})
	}

	return c.JSON(http.StatusOK, result)
}
