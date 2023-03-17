package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	myMsg := []byte("hiiiiii")
	//	c1 := http.Cookie{
	//		Name:     "Hello1",
	//		Value:    base64.URLEncoding.EncodeToString(myMsg),
	//		HttpOnly: true,
	//	}
	c1 := http.Cookie{
		Name:  "Gopher",
		Value: base64.URLEncoding.EncodeToString(myMsg),
	}

	c2 := http.Cookie{
		Name:  "Lama",
		Value: base64.URLEncoding.EncodeToString(myMsg),
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	fmt.Fprint(w, "this is the set handler")
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	// this checks for a Name value of "Hello Message"
	theCookieValue, err := r.Cookie("Lama")
	if err != nil {
		// if the err is not nill (there is an errot)
		// checks if it falls into the ErrNoCookie instance and if it does
		// print the message to the client.
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No such cookie.")
		}
		// updates the cookie if the cookie with the Name "Other Message" is present.
	} else {
		rc := http.Cookie{
			Name:    "Lama",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(theCookieValue.Value)
		fmt.Fprintln(w, string(val))

	}
	fmt.Fprint(w, "This is the getmessage handler")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/show", showMessage)

	fmt.Println("server is starging....")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("the server cannot start", err)
	}
}
