package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type Myhandler struct {}

func chainGang(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Chain Chain"))
}

func hello2(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Yelllllllooooow"))
}

// log -> creartes a log that tell the developer that a handler function with the name
    // name is being requested to the client using closure.
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function name is called: ", name)
		h(w, r)
	}
}

type helloHand struct {}

func (h *helloHand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A new Hello"))
}

type worldHand struct {}

func (h *worldHand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A new world"))
}
func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloooooo"))
}
func main() {
	fmt.Printf("Server statring....\n")


	hello := helloHand{}
	world := worldHand{}
	// no url matchung needed so any path will be a valid request.
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello2", log(hello2))
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	// starts the server listens for request on a specific addy in this case 8080 local host
	server.ListenAndServe()
}
