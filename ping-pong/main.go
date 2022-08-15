package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var count int

	app := gin.Default()

	app.GET("/pingpong", func(c *gin.Context) {
		c.Writer.Write([]byte(fmt.Sprintln("pong", count)))
		writeToFile(fmt.Sprint(count))

		count++
	})

	run(app, "9090")

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}

func writeToFile(data string) {

	file, err := os.OpenFile("/tmp/pingpong.log", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err = file.WriteString(data); err != nil {
		fmt.Println(err)
	}

}
