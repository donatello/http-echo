package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	for k, v := range r.Header {
		for _, val := range v {
			fmt.Printf("%s: %s\n", k, val)
		}
	}
	fmt.Println("")

	defer r.Body.Close()
	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Printf("Error writing to stdout: %v", err)
	}
	fmt.Println("")
}

func main() {
	http.HandleFunc("/", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
