package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var count int

	app := gin.Default()

	app.GET("/pingpong", func(c *gin.Context) {
		c.Writer.Write([]byte(fmt.Sprintln("pong", count)))

		count++
	})

	run(app, "9090")

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}
