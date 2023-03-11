package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
    // create a message of type bytes
    msg := []byte("Hello golang")
    // creates an instance of the cookie for the client to recieve
    cookieMsg := http.Cookie{
        Name: "Hello Message",
        // the Value is
        Value: base64.URLEncoding.EncodeToString(msg),
        HttpOnly: true,
        }
        http.SetCookie(w, &cookieMsg)
    fmt.Fprint(w, "this is the setMessage handler")
}

func newCookie(w http.ResponseWriter, r *http.Request)  {
    myMsg := []byte("hiiiiiiii")
    c1 := http.Cookie{
        Name: "Hello",
        Value: base64.URLEncoding.EncodeToString(myMsg),
        HttpOnly: true,
    }
    http.SetCookie(w, &c1)
    fmt.Fprint(w, "this is the new handler")
}

func getMessage(w http.ResponseWriter, r *http.Request) {
    // this checks for a Name value of "Hello Message"
    theCookieValue, err := r.Cookie("Hello Message")
    if err != nil {
        // if the err is not nill (there is an errot)
        // checks if it falls into the ErrNoCookie instance and if it does
        // print the message to the client.
        if err == http.ErrNoCookie {
            fmt.Fprintln(w, "No such cookie.")
        }
    } else {
        rc := http.Cookie{
            Name:    "Hello Message",
            Value:   theCookieValue.Value,
            MaxAge:  3600,
            Expires: time.Now().Add(24 * time.Hour),
            }

            http.SetCookie(w, &rc)
        fmt.Fprintln(w, theCookieValue.Value)
    }
    fmt.Fprint(w, "This is the getmessage handler")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
    http.HandleFunc("/set", setMessage)
    http.HandleFunc("/new", newCookie)

    http.HandleFunc("/get_message", getMessage)
	fmt.Println("server is starging....")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("the server cannot start", err)
	}
}