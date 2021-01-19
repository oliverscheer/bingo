package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const port = 8080

func main() {
	http.HandleFunc("/catalog/products", productsHandler)

	http.HandleFunc("/alive", aliveHandler)

	http.HandleFunc("/helloworld", HelloWorld)

	// start
	addr := ":" + strconv.Itoa(port)
	fmt.Println("Start listening on port", addr)
	http.ListenAndServe(addr, nil)
}

// HelloWorld is just a test function
func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Schuhe, Hose, Hemd"))
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm alive"))
}
