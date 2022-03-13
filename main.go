package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/blockchain"
	"github.com/daveg7lee/nomadcoin/utils"
)

const port string = ":4000"

type URL string

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func createDocs() []URLDescription {
	return []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "GET",
			Description: "Get all blocks",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See a block",
		},
	}
}

func handleDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	docs := createDocs()
	json.NewEncoder(w).Encode(docs)
}

func handleBlocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBlocks(w)
	case "POST":
		postBlocks(w, r)
	}
}

func getBlocks(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain.GetBlockchain().GetAllBlocks())
}

func postBlocks(w http.ResponseWriter, r *http.Request) {
	var addBlockBody AddBlockBody
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
	blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	fmt.Printf("Server listening on http://localhost%s\n", port)
	http.HandleFunc("/", handleDocs)
	http.HandleFunc("/blocks", handleBlocks)
	log.Fatal(http.ListenAndServe(port, nil))
}
