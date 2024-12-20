package main

import (
	"fmt"
	"log"

	"github.com/otevchik/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it works", file)

}
