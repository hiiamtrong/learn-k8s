package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	randomStr, err := uuid.NewUUID()

	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		hash := readFile("/tmp/writer.log")
		pingCount := readFile("/tmp/pingpong.log")
		c.JSON(200, fmt.Sprintf("%s: %s\n Ping / Pong: %s", randomStr, hash, pingCount))

	})

	run(app, "8080")

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}

func readFile(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	return string(data)
}
