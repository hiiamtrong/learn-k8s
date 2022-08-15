package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	var timeStr string

	for {
		timeStr = time.Now().Format("2006-01-02T15:04:05Z07:00")
		fmt.Printf("%s\n", timeStr)
		writeToFile(fmt.Sprint(timeStr))
		time.Sleep(time.Second * 5)
	}

}

func writeToFile(data string) {

	file, err := os.OpenFile("/tmp/writer.log", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err = file.WriteString(data); err != nil {
		fmt.Println(err)
	}

}
