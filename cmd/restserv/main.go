package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Healthy",
	}
	js, _ := json.Marshal(response)
	_, _ = w.Write(js)

}
func main() {
	http.HandleFunc("/health-check", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
