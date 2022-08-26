package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/ping-ping/config"
	"github.com/hiiamtrong/ping-ping/database"
	"github.com/hiiamtrong/ping-ping/model"
)

func main() {
	config.Init()
	database.Init()

	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})
	app.GET("/pingpong", func(c *gin.Context) {
		counter := getLatestCounter()
		count := -1
		if counter != nil {
			count = int(counter.Count)
		}

		count++
		c.Writer.Write([]byte(fmt.Sprintln("pong", count)))
		saveToDb((count))

	})

	app.GET("/pingpong/count", func(c *gin.Context) {
		counter := getLatestCounter()
		count := 0
		if counter != nil {
			count = int(counter.Count)
		}

		c.Writer.Write([]byte(fmt.Sprint(count)))

	})

	log.Fatal(run(app, "9090"))

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}

func saveToDb(data int) {
	_, err := database.DB.Exec("INSERT INTO counter (count) VALUES ($1)", data)
	if err != nil {
		fmt.Println(err)
	}

}

func getLatestCounter() *model.Counter {
	res := database.DB.QueryRow("SELECT * FROM counter ORDER BY id DESC LIMIT 1 ")

	if res == nil {
		return nil
	}

	var (
		id    int64
		count int64
	)

	if err := res.Scan(&id, &count); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &model.Counter{
		ID:    id,
		Count: id,
	}
}
