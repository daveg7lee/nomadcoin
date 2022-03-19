package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type errorResponse struct {
	ErrorMessage string `json:"error message"`
}

type balanceResponse struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

type addTxPayload struct {
	To     string
	Amount int
}

func createDocs() []document {
	return []document{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
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
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
		{
			URL:         url("/mempool"),
			Method:      "GET",
			Description: "Get mempool",
		},
		{
			URL:         url("/transaction"),
			Method:      "POST",
			Description: "Make a Transaction",
		},
	}
}

func handleDocs(w http.ResponseWriter, r *http.Request) {
	docs := createDocs()
	json.NewEncoder(w).Encode(docs)
}

func handleBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blockchain().Blocks())
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
	data, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(w)
	if err == blockchain.ErrorNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(data)
	}
}

func postBlock(w http.ResponseWriter, r *http.Request) {
	blockchain.Blockchain().AddBlock()
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

func handleStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blockchain())
}

func handleBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")

	switch total {
	case "true":
		handleTotalBalance(w, address)
	default:
		handleTxOuts(w, address)
	}
}

func handleTotalBalance(w http.ResponseWriter, address string) {
	amount := blockchain.Blockchain().BalanceByAddress(address)
	utils.HandleErr(json.NewEncoder(w).Encode(balanceResponse{
		Address: address,
		Amount:  amount,
	}))
}

func handleTxOuts(w http.ResponseWriter, address string) {
	txOuts := blockchain.Blockchain().TxOutsByAddress(address)
	utils.HandleErr(json.NewEncoder(w).Encode(txOuts))
}

func handleMempool(w http.ResponseWriter, r *http.Request) {
	mempool := blockchain.Mempool
	utils.HandleErr(json.NewEncoder(w).Encode(mempool.Txs))
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{"no enough money"})
	}
	w.WriteHeader(http.StatusCreated)

}

func Start(portNum int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", portNum)

	router.Use(setJsonContentTypeMiddleware)
	router.HandleFunc("/", handleDocs).Methods("GET")
	router.HandleFunc("/status", handleStatus).Methods("GET")
	router.HandleFunc("/blocks", handleBlocks).Methods("GET")
	router.HandleFunc("/block", handleBlock).Methods("POST")
	router.HandleFunc("/block/{hash:[a-f0-9]+}", handleBlock).Methods("GET")
	router.HandleFunc("/balance/{address}", handleBalance).Methods("GET")
	router.HandleFunc("/mempool", handleMempool).Methods("GET")
	router.HandleFunc("/transaction", handleTransactions).Methods("POST")

	fmt.Printf("Rest Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
