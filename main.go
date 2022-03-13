package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func createDocs() []URLDescription {
	return []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/block",
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	docs := createDocs()
	json.NewEncoder(w).Encode(docs)
}

func main() {
	fmt.Printf("Server listening on http://localhost%s\n", port)
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(port, nil))
}
