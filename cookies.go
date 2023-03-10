package main

import (
	"net/http"
)

func setCookies(writer http.ResponseWriter, req *http.Request) {
	// this creates a new instance of Cookies with 3 fields set
	c1 := http.Cookie {
		Name: "first cookie",
		Value: "Go web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie {
		Name: "Second Cookie",
		Value: "Something else",
		HttpOnly: true,
	}
	http.SetCookie(writer, &c1)
	http.SetCookie(writer, &c2)

}
func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set-cookies", setCookies)
	server.ListenAndServe()
	
}