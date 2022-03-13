package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/nomadcoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	bytes, err := json.Marshal(data)
	utils.HandleErr(err)

}

func main() {
	fmt.Printf("Server listening on http://localhost%s\n", port)
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(port, nil))
}
