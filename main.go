package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = "2024"

func main() {
	http.HandleFunc("/debug", debugHandler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := r.URL.Query()["num"]
	if !ok || len(data) < 0 {
		log.Println("param is missing")
		return
	}

	number, err := strconv.Atoi(data[0])
	if err != nil {
		log.Printf("invalid parameter: '%v'", err)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("input number test: %v", number)))
	if err != nil {
		log.Printf("failed to return")

		return
	}
}
