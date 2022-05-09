package main

import (
	"final/cmd"
	"final/cmd/router"
	"log"
	"net/http"
)

func main() {
	router := router.New()
	// router.Start(":3000")

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}

// tasks = append(tasks, Task{Id: 10, Text: "cucumber", ListId: 1, Completed: true})
// tasks = append(tasks, Task{Id: 9, Text: "bread", ListId: 1, Completed: true})
// tasks = append(tasks, Task{Id: 7, Text: "milk", ListId: 1, Completed: true})
// tasks = append(tasks, Task{Id: 6, Text: "jogging", ListId: 2, Completed: true})
// tasks = append(tasks, Task{Id: 5, Text: "soap", ListId: 1, Completed: true})
// tasks = append(tasks, Task{Id: 4, Text: "basketball", ListId: 2, Completed: true})
// tasks = append(tasks, Task{Id: 3, Text: "fishing", ListId: 3, Completed: true})
// tasks = append(tasks, Task{Id: 2, Text: "reading", ListId: 3, Completed: true})
// tasks = append(tasks, Task{Id: 1, Text: "table tennis", ListId: 2, Completed: true})

// lists = append(lists, List{Id: 1, Name: "groceries"})
// lists = append(lists, List{Id: 2, Name: "sports"})
// lists = append(lists, List{Id: 3, Name: "hobbies"})
