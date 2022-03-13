package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/blockchain"
	"github.com/daveg7lee/nomadcoin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type document struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func createDocs() []document {
	return []document{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "Get all blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
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
	var addBlockBody addBlockBody

	utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
	blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
	w.WriteHeader(http.StatusCreated)
}

func Start(portNum int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", portNum)

	handler.HandleFunc("/", handleDocs)
	handler.HandleFunc("/blocks", handleBlocks)

	fmt.Printf("Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
