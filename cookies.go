package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	// creates a instance of sessions cookie
	c1 := http.Cookie{
		Name:     "cookie1",
		Value:    "My cookie",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "cookie2",
		Value:    "second cookie value",
		HttpOnly: true,
	}

	// this adds the cookies to the response headers
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("the funciton %s has been called\n", name)
		// this will call the handler function that is inside of the
		// log function
		h(w, r)
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("cookie1")
	if err != nil {
		fmt.Fprintln(w, "cookie can not be found")
	}
	cs := r.Cookies()
	// this will display the cookies in
	fmt.Fprintf(w, "this is the c1 %s\n", c1)
	// this will display the cookies in the cs variable
	fmt.Fprintln(w, cs)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/setCookie", log(setCookie))
	http.HandleFunc("/getCookie", log(getCookie))
	fmt.Println("server .....")
	server.ListenAndServe()
}
