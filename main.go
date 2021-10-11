package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"githun.com/heuristicwave/buildingBlockchain/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string // 대,소문자 (public, private)
	Blocks    []*blockchain.Block
}

// homeHandler
func home(rw http.ResponseWriter, r *http.Request) {
	// ParseFiles 는 에러를 return 하므로 같이 지정하는데,
	// Must를 사용하여 if err!=nil 까지 대체
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	// Template에 Handler에서 만든 data 전달
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home) // route
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
