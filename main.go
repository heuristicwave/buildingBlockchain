package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// homeHandler
func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from Server!") // write on Response
}

func main() {
	http.HandleFunc("/", home) // route
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
