package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/todo-simple/config"
	"github.com/hiiamtrong/todo-simple/db"
	"github.com/hiiamtrong/todo-simple/models"
	"github.com/hiiamtrong/todo-simple/repository"
)

var todoRepo = new(repository.TodoRepo)

func main() {

	gin.SetMode(gin.ReleaseMode)

	config.Init()
	db.Init()

	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	app.Use(cors.New(config))

	app.GET("/api", func(c *gin.Context) {

		todos, err := todoRepo.GetTodos()

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, todos)
	})

	app.POST("/api", func(c *gin.Context) {
		todo := models.Todo{}
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := todoRepo.AddTodo(&todo)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

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
