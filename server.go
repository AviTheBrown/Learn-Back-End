package main

import (
	"fmt"
	"net/http"
)
func handler(writer http.ResponseWriter, request *http.Request) {
    // will write back to the client -> localhost:8080 using the writter argument
//    fmt.Fprintf(writer, "hellooooooooo world,")
    fmt.Fprintf(writer, "Hello world")
}
func main() {
    fmt.Println("Server is running...")
    // creates a HTTP server with .HandleFunc
        // the URL path "/" means that this the "handler" function
        // will be called for any incoming request that doesnt match a more specific URL path.
            // ex google.com/images -> in this case the path is "images" the domain name is "google.com"
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
