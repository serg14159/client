package main

import (
	"fmt"
	"log"

	"github.com/serg14159/storage/pkg/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client", file)
}
