package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	randomStr, err := uuid.NewUUID()
	var timeStr string
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			timeStr = time.Now().Format("2006-01-02T15:04:05Z07:00")
			fmt.Printf("%s: %s\n", timeStr, randomStr)
			time.Sleep(time.Second * 5)
		}
	}()

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"string": randomStr,
			"time":   timeStr,
		})
	})

	run(app, "8080")

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}
