package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := os.Args[1]
	if port == "" {
		port = "8000"
	}
	fmt.Println("Serving files in the current directory on port " + port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
