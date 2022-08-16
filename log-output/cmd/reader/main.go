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
		res, err := http.Get("http://ping-pong-svc:2223/pingpong/count")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, fmt.Sprintf("%s: %s\n Ping / Pong: %s", randomStr, hash, string(data)))

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
