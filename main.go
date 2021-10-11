package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"githun.com/heuristicwave/buildingBlockchain/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string // 대,소문자 (public, private)
	Blocks    []*blockchain.Block
}

// homeHandler
func home(rw http.ResponseWriter, r *http.Request) {
	// Template에 Handler에서 만든 data 전달
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home) // route
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
