package main

import (
	"fmt"
	"log"
	"os"
)

func fileWrite() {
	fmt.Println("<file write>")
	err := os.WriteFile("bob.md", []byte("test"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
