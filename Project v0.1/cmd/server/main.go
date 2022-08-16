package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/todo-simple/config"
)

type Todo struct {
	Title string `json:"title" binding:"required"`
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	config.Init()

	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	app.Use(cors.New(config))

	todos := []Todo{}

	app.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	app.POST("/api", func(c *gin.Context) {
		todo := Todo{}
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todos = append(todos, todo)
		c.JSON(http.StatusOK, gin.H{
			"message": "Todo added",
		})
	})

	log.Fatal(run(app))
}

func run(app *gin.Engine) error {
	log.Println("Server started on port", config.Cfg.PortServer)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Cfg.PortServer), app.Handler())
}
