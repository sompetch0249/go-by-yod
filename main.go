package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hello/go-by-yod/fizzbuzz"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.Say(n))
}
