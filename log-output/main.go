package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {

	randomStr, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02T15:04:05Z07:00"), randomStr)
		time.Sleep(time.Second * 5)
	}
}
