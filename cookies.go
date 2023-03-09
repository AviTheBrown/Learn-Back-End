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
	// this creates Header text of "Set-Cookie" to the stringify version of c1 and c2
	writer.Header().Set("Set-Cookie", c1.String())
	writer.Header().Set("Set-Cookie", c2.String())

}
func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	
}