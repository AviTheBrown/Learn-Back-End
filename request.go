package main

import (
	"fmt"
	"net/http"
)

func headers(writer http.ResponseWriter, req *http.Request) {
    // this retreives information from the Head of the request
    // there should be no whitespace in the string arg
	head := req.Header["Accept-Encoding"]
    // or
    head2 := req.Header.Get("Accept-Encoding")

    // prints a map of string
    fmt.Fprintln(writer, head)
    // prints a comma-delimited list of values
    writer.Write([]byte(head2))
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", headers)
    fmt.Printf("Starting server.....")
	server.ListenAndServe()

}
