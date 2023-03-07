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

func body(writer http.ResponseWriter, req *http.Request)  {
    // creates a var that represents the length of the message body in bytes,
    // not the header of the request.
    len := req.ContentLength
    // creates a slice of bytes that is equal to the length of the body message
    // of the request
    body := make([]byte, len)
    // reads all the body message into the body slice
    req.Body.Read(body)
    defer req.Body.Close()
    writer.Write([]byte(body))
}

func process(writer http.ResponseWriter, req *http.Request)  {
    req.ParseForm()
    fmt.Fprintln(writer, req.Form)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
//	http.HandleFunc("/header", headers)
//    http.HandleFunc("/body", body)
    http.HandleFunc("/process", process)
    fmt.Printf("Starting server.....")
	server.ListenAndServe()

}
