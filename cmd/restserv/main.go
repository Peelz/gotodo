package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var port string

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := map[string]interface{}{
		"message": "Healthy",
	}
	js, _ := json.Marshal(response)
	_, _ = w.Write(js)
}

func init() {
	defer flag.Parse()
	flag.StringVar(&port, "port", "80", "Service port")
}

func main() {

	http.HandleFunc("/health-check", healthCheck)
	log.Printf("Listing on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
