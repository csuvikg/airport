package main

import (
	"airport/cmd/api"
	"net/http"
)

func main() {
	http.HandleFunc("/list", api.ListAirports)
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	_ = http.ListenAndServe(":3003", nil)
}
