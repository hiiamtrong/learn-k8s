package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hiiamtrong/todo-simple/config"
	"github.com/hiiamtrong/todo-simple/db"
	"github.com/hiiamtrong/todo-simple/models"
	"github.com/hiiamtrong/todo-simple/repository"
)

func main() {

	config.Init()
	db.Init()
	log.Println("Starting adding random article...")
	var todoRepo = new(repository.TodoRepo)

	res, err := http.Get("http://en.wikipedia.org/w/api.php?action=query&list=random&format=json&rnnamespace=0&rnlimit=1")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Get random article error with status %d", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	resData := make(map[string]interface{})

	err = json.Unmarshal(data, &resData)

	if err != nil {
		log.Fatal(err)
	}
	title := (resData["query"].(map[string]interface{})["random"].([]interface{})[0].(map[string]interface{})["title"])
	todoTitle := fmt.Sprintf("Read https://en.wikipedia.org/wiki/%s", title)
	err = todoRepo.AddTodo(
		&models.Todo{
			Title: todoTitle,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Added todo: %s", todoTitle)

}
