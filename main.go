package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from home!")
}

func main() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(port, nil))
}
