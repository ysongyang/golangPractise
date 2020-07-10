package main

import (
	"fmt"
	"github.com/gitstliu/go-id-worker"
	uuid "github.com/satori/go.uuid"
)

func main() {
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newId, _ := currWoker.NextId()
	fmt.Println(newId)

	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)
}
