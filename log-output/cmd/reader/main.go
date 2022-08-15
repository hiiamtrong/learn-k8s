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
		data := readFile()
		c.JSON(200, fmt.Sprintf("%s: %s", randomStr, data))

	})

	run(app, "8080")

}

func run(app *gin.Engine, port string) error {
	log.Println("Server started on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler())
}

func readFile() string {
	file, err := os.OpenFile("/tmp/writer.log", os.O_RDONLY, 0644)

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
