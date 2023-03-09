package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type POST struct {
    User string
    Threads []string
}

func writeExmaple (writer http.ResponseWriter,  req *http.Request) {
    str := `<html>
    <head><title> Go Web </title></head>
        <body><h1> Hello World </h1></body>
        </html>`

    writer.Write([]byte(str))
}

func writerHeaderExample(writer http.ResponseWriter, req *http.Request) {
    writer.WriteHeader(501)
    fmt.Fprintln(writer, "No such servixe.")
}

func headerExample(writer http.ResponseWriter, req *http.Request)  {
    writer.Header().Set("location", "http://netflix.com")
    writer.WriteHeader(302)
}

func jsoneExample(writer http.ResponseWriter, req *http.Request)  {
    writer.Header().Set("Content-Tye", "application/json")
    post := POST{
        User: "Avi Brown",
        Threads: []string{"first", "second", "third"},
    }
    json, _ := json.Marshal(post)
    writer.Write(json)
}
func main() {

    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/write", writeExmaple)
    http.HandleFunc("/writeHeader", writerHeaderExample)
    http.HandleFunc("/header", headerExample)
    http.HandleFunc("/json", jsoneExample)
    fmt.Println("Server starting")
    server.ListenAndServe()

}