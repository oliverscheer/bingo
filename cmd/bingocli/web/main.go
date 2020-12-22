package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const port = 8080

func main() {
	http.HandleFunc("/catalog/products", productsHandler)

	// start
	addr := ":" + strconv.Itoa(port)
	fmt.Println("Start listening on port", addr)
	http.ListenAndServe(addr, nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Schuhe, Hose, Hemd"))
}
