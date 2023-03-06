package main

import (
	"fmt"
	"net/http"
)

//// *** the client is the web browser ***
//func handler(writer http.ResponseWriter, request *http.Request) {
//	// will write back to the client -> browser (localhost:8080) using the ResponWriter argument from thr net/http package
//	//    fmt.Fprintf(writer, "hellooooooooo world,")
//	fmt.Fprintf(writer, "Hello world")
//}

//func main() {
//    fmt.Println("Server is running...")
//    // creates a HTTP server with .HandleFunc
//        // the URL path "/" means that this the "handler" function
//        // will be called for any incoming request that doesnt match a more specific URL path.
//            // ex google.com/images -> in this case the path is "images" the domain name is "google.com"
//	http.HandleFunc("/", handler)
//    // sets up the HTTP SERVER to listen on Port 8080 of the Local Machine -> localhost:
//        // "8080" = the network address, nil = DeafultServeMux this is the default multiplexer.
//	http.ListenAndServe(":8080", nil)
//}

// creates a
type HelloHandler struct{}

// method for Hello struct
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type WorldHandler struct{}

// method for the World struct
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	fmt.Println("Server staring....")
    // create instace of the two structs
	hello := HelloHandler{}
	world := WorldHandler{}

    // create the server
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

    // handle the request in according to the URL path and calling the instance of that handler.
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

    // start server
	server.ListenAndServe()

}
