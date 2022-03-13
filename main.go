package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/blockchain"
	"github.com/daveg7lee/nomadcoin/utils"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*block.Block
	Year      int
}

type addData struct {
	PageTitle string
	Year      int
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := homeData{PageTitle: "Home", Blocks: blockchain.GetBlockchain().GetAllBlocks(), Year: utils.GetYear()}
	templates.ExecuteTemplate(w, "home", data)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAdd(w)
	case "POST":
		postAdd(w, r)
	}
}

func getAdd(w http.ResponseWriter) {
	data := addData{PageTitle: "Add", Year: utils.GetYear()}
	templates.ExecuteTemplate(w, "add", data)
}

func postAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := r.Form.Get("blockData")
	blockchain.GetBlockchain().AddBlock(data)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func main() {
	fmt.Printf("Server running on http://localhost%s\n", port)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	log.Fatal(http.ListenAndServe(port, nil))
}
