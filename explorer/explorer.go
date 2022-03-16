package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/daveg7lee/nomadcoin/block"
	"github.com/daveg7lee/nomadcoin/blockchain"
	"github.com/gorilla/mux"
)

const templateDir string = "explorer/templates/"

var templates *template.Template

type pageData struct {
	PageTitle string
	Blocks    []*block.Block
	Year      int
}

func getYear() int {
	now := time.Now()
	return now.Year()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	//data := pageData{PageTitle: "Home", Blocks: blockchain.Blockchain().AllBlocks(), Year: getYear()}
	//templates.ExecuteTemplate(w, "home", data)
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
	data := pageData{PageTitle: "Add", Year: getYear()}
	templates.ExecuteTemplate(w, "add", data)
}

func postAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := r.Form.Get("blockData")
	blockchain.Blockchain().AddBlock(data)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func loadTemplates() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
}

func Start(portNum int) {
	router := mux.NewRouter()
	port := fmt.Sprintf(":%d", portNum)

	loadTemplates()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/add", handleAdd).Methods("GET, POST")

	fmt.Printf("Explorer Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
