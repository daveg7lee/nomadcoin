package main

import (
	"fmt"
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
	html := template.Must(template.ParseFiles("./templates/pages/home.gohtml"))
	data := homeData{PageTitle: "Home", Blocks: blockchain.GetBlockchain().GetAllBlocks()}
	html.Execute(w, data)
}

func main() {
	fmt.Printf("Server running on http://localhost%s\n", port)
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(port, nil))
}
