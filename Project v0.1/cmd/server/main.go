package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/todo-simple/config"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	config.Init()

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	log.Fatal(run(app))
}

func run(app *gin.Engine) error {
	log.Println("Server started on port", config.Cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Cfg.Port), app.Handler())
}
