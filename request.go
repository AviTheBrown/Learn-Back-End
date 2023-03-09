package main

import (
	"fmt"
	"net/http"
    "reflect"
    "runtime"
)

func headers(writer http.ResponseWriter, req *http.Request) {
    // this retreives information from the Head of the request
    // there should be no whitespace in the string arg
    fmt.Fprintf(writer, "this is the header\n")
//	head := req.Header["Accept-Encoding"]
    // or
    head2 := req.Header.Get("Accept-Language")

    // prints a map of string
//    fmt.Fprintln(writer, head)
    // prints a comma-delimited list of values
    writer.Write([]byte(head2))
}

func body(writer http.ResponseWriter, req *http.Request)  {
    fmt.Fprintf(writer, "this is the body\n")

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
func log(h http.HandlerFunc) http.HandlerFunc {
    return func(writer http.ResponseWriter, req *http.Request) {
        name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
        fmt.Println( "The handler %T was invoked\n", name)
        h(writer, req)
    }
}
func process(writer http.ResponseWriter, req *http.Request)  {
    fmt.Fprintf(writer, "this is the process\n")
    // parses the form from the clienr
    req.ParseForm()
    // parses form usinf a specified amount of memory to perform the action
    //
    req.ParseMultipartForm(1024)
    // this only parses and retireves the key-value pair of the form and omits the
    // URL key-value pairs.
    fmt.Fprintln(writer, req.PostForm)

    // prints both the URL and form KV pairs
    fmt.Fprintln(writer, req.Form)

    // MultiparseForm only contains the form KV pairs
    // only with "multipart/form-data" enctype
    // if you use application/x-www-form-urlencoded enctype it will return nil
    fmt.Fprintln(writer, req.MultipartForm)

    // access the KV pair from the form without have to parse the form
    // it does it by itself
    // with "application/x-www-form-urlencoded" :
    // it retrieves the first value associated with the given key from the request body,
    // and it ignores other values with the same key.
    // --------------
    // with  "multipart/form-data"
    // it  retrieves the first value associated with the given key from the request body,
    // and it ignores other values with the same key.
    // However, it also looks for the value in the URL query parameters
    // and the path parameters, and it returns the first value found.
    fmt.Fprintf(writer, req.FormValue("hello"))
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", headers)
    http.HandleFunc("/body", body)
    http.HandleFunc("/process", log(process))
    fmt.Printf("Starting server.....\n")
	server.ListenAndServe()
}
