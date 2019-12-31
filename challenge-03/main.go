package main

import (
	"encoding/json"
	"github.com/anthonyrichir/go-challenges/challenge-03/stats"
	"net/http"
	"log"
)

func handleStats(w http.ResponseWriter, r *http.Request) {
	s := stats.FetchStats()
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/stats", handleStats)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
