package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Println("Serving files in the current directory on port " + port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
