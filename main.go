package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-by-yod/fizzbuzz"
	"github.com/go-by-yod/oscar"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	r.HandleFunc("/oscarlook", oscarHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.Say(n))
}

func oscarHandler(w http.ResponseWriter, req *http.Request) {
	m := oscar.ActorWhoWonMoreThanOneAward("./oscar/oscar_age_male.csv")

	w.Header().Set("Content-type", "text/json")
	json.NewEncoder(w).Encode(&m)
}
