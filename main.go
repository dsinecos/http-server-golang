package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Test-Header", "Value for Test Header")
	fmt.Fprintln(w, "Response from Hotdog handler")
	fmt.Printf("Method %s \n", req.Method)
	fmt.Printf("User-Agent %s \n", req.UserAgent())
}

func main() {
	var h hotdog
	err := http.ListenAndServe(":9080", h)

	if err != nil {
		log.Fatalln(err)
	}

}
