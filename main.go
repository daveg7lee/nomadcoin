package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*block.Block
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := homeData{PageTitle: "Home", Blocks: blockchain.GetBlockchain().GetAllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "add", nil)
}

func main() {
	fmt.Printf("Server running on http://localhost%s\n", port)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	log.Fatal(http.ListenAndServe(port, nil))
}
