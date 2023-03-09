package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Does this work..."))
}

func hello2(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("hello2"))
}
//func log(h http.HandlerFunc) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
//		fmt.Printf("The handler %T was called\n", name)
//		// the return handlerFunc
//		h(writer, request)
//	}
//}
// testing
func main() {
	hello := hello2

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	fmt.Printf("starting server....")
	http.HandleFunc("/hello2", log(hello))
	server.ListenAndServe()
}
