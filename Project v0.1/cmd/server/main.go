package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/todo-simple/config"
)

func main() {

	go func() {
		if !fileExists("./public/thumbnail.png") {
			err := saveImageFromUrl("https://picsum.photos/1200")

			if err != nil {
				log.Println(err)
			} else {
				log.Println("Create thumbnail.png")
			}

		} else {
			for {

				if time.Now().Hour() == 0 && time.Now().Minute() == 0 {
					err := saveImageFromUrl("https://picsum.photos/1200")
					if err != nil {
						log.Println(err)
					} else {
						log.Println("Update thumbnail.png")
					}
				}
			}
		}

	}()

	gin.SetMode(gin.ReleaseMode)

	config.Init()

	app := gin.Default()

	app.LoadHTMLGlob("templates/*")
	app.Static("/public", "./public")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Fatal(run(app))
}

func run(app *gin.Engine) error {
	log.Println("Server started on port", config.Cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Cfg.Port), app.Handler())
}

func saveImageFromUrl(url string) error {
	// TODO
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	file, err := os.OpenFile("./public/thumbnail.png", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}
	return nil

}

// check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
