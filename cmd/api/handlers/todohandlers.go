package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

var tasks []Task
var lists []List

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

func getNextTaskID() int {
	highestID := -1
	if len(tasks) == 0 {
		return 1
	}
	for _, task := range tasks {
		if highestID < task.Id {
			highestID = task.Id
		}
	}
	return highestID + 1
}

func getNextListID() int {
	highestID := -1
	if len(lists) == 0 {
		return 1
	}
	for _, list := range lists {
		if highestID < list.Id {
			highestID = list.Id
		}
	}
	return highestID + 1
}

func DeleteList(c echo.Context) error {
	str := strings.Split(c.Request().URL.Path, "/")
	idPath := str[2]
	id, err := strconv.Atoi(idPath)
	if err != nil {
		log.Printf("Failed to convert id in deleteList: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	for i := 0; i < len(tasks); i++ {
		if id == tasks[i].ListId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			i--
		}
	}

	for i, list := range lists {
		if id == list.Id {
			lists = append(lists[:i], lists[i+1:]...)
			return c.String(http.StatusOK, "")
		}
	}

	return c.String(http.StatusInternalServerError, "")
}

func DeleteTask(c echo.Context) error {
	str := strings.Split(c.Request().URL.Path, "/")
	idPath := str[2]
	id, err := strconv.Atoi(idPath)
	if err != nil {
		log.Printf("Failed to convert id in deleteTask: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	for i, task := range tasks {
		if id == task.Id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.String(http.StatusOK, "")

		}
	}
	return c.String(http.StatusInternalServerError, "")
}

func ToggleTask(c echo.Context) error {
	str := strings.Split(c.Request().URL.Path, "/")
	idPath := str[2]
	id, err := strconv.Atoi(idPath)
	if err != nil {
		log.Printf("Failed to convert id: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	for i, task := range tasks {
		if id == task.Id {
			task.Completed = true
			tasks = append(tasks[:i], tasks[i+1:]...)
			tasks = append(tasks[:i+1], tasks[i:]...)
			tasks[i] = task
			return c.JSON(http.StatusOK, task)
		}
	}
	return c.JSON(http.StatusInternalServerError, "")
}

func AddList(c echo.Context) error {
	l := List{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &l)
	if err != nil {
		log.Printf("Failed unmarshaling in addList: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if l.Id != 0 {
		log.Printf("Filling in ID is prohibited.")
	}

	id := getNextListID()
	l.Id = id

	lists = append(lists, l)
	return c.JSON(http.StatusOK, l)
}

func AddTask(c echo.Context) error {
	str := strings.Split(c.Request().URL.Path, "/")
	idPath := str[2]
	listId, err := strconv.Atoi(idPath)
	if err != nil {
		log.Printf("Failed to convert id: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}
	t := Task{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addTask: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		log.Printf("Failed unmarshaling in addTask: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if t.Id != 0 || t.ListId != 0 {
		log.Println("Filling in IDs is prohibited.")
	}

	var hasList bool

	for _, l := range lists {
		if listId == l.Id {
			hasList = true
		}
	}

	if !hasList {
		log.Printf("List with ID %d has not been created yet.\n", listId)
		return c.String(http.StatusInternalServerError, "")
	}

	id := getNextTaskID()
	t.Id = id
	t.ListId = listId

	tasks = append(tasks, t)
	return c.JSON(http.StatusOK, t)
}

func GetTasks(c echo.Context) error {
	str := strings.Split(c.Request().URL.Path, "/")
	idPath := str[2]
	id, err := strconv.Atoi(idPath)
	if err != nil {
		log.Printf("Failed to convert id: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}
	var result []Task
	for _, task := range tasks {
		if id == task.ListId {
			result = append(result, task)
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetLists(c echo.Context) error {
	return c.JSON(http.StatusOK, lists)
}
