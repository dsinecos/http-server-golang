package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Response from Hotdog handler")
}

func main() {
	var h hotdog
	err := http.ListenAndServe(":9080", h)

	if err != nil {
		log.Fatalln(err)
	}

}
