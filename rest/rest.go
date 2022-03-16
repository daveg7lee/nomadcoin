package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/blockchain"
	"github.com/daveg7lee/nomadcoin/utils"
	"github.com/gorilla/mux"
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

type errorResponse struct {
	ErrorMessage string `json:"error message"`
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
			URL:         url("/block"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/block/{hash}"),
			Method:      "GET",
			Description: "See a block",
		},
	}
}

func handleDocs(w http.ResponseWriter, r *http.Request) {
	docs := createDocs()
	json.NewEncoder(w).Encode(docs)
}

func handleBlocks(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(blockchain.Blockchain().AllBlocks())
}

func handleBlock(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBlock(w, r)
	case "POST":
		postBlock(w, r)
	}
}

func getBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	data, err := block.FindBlock(hash)
	encoder := json.NewEncoder(w)
	if err == block.ErrorNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(data)
	}
}

func postBlock(w http.ResponseWriter, r *http.Request) {
	var addBlockBody addBlockBody

	utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
	blockchain.Blockchain().AddBlock(addBlockBody.Message)
	w.WriteHeader(http.StatusCreated)
}

func setJsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func Start(portNum int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", portNum)

	router.Use(setJsonContentTypeMiddleware)
	router.HandleFunc("/", handleDocs).Methods("GET")
	router.HandleFunc("/blocks", handleBlocks).Methods("GET")
	router.HandleFunc("/block", handleBlock).Methods("POST")
	router.HandleFunc("/block/{hash:[a-f0-9]+}", handleBlock).Methods("GET")

	fmt.Printf("Rest Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
