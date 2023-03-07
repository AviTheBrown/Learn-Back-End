package main

import (
    "fmt"
    "net/http"
)

type Myhandler struct {
}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("helloooooo"))
}
func main() {
    fmt.Print("Server statring....")
    handler := Myhandler{}
    // no url matchung needed so any path will be a valid request.
    server := http.Server{
        Addr: "127.0.0.1:8080",
        Handler: &handler,
    }
    // starts the server listens for request on a specific addy in this case 8080 local host
    server.ListenAndServe()
}