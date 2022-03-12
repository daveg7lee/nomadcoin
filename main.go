package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*block.Block
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("./templates/home.html"))
	data := homeData{PageTitle: "Home", Blocks: blockchain.GetBlockchain().GetAllBlocks()}
	html.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(port, nil))
}
