package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Println("Stoa läuft auf http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
