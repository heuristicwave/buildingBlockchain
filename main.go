package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"githun.com/heuristicwave/buildingBlockchain/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

// Sending JSON our users
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	jsonByte, err := json.Marshal(data) // (byte, error)
	utils.HandleErr(err)
	fmt.Printf("%s", jsonByte)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
