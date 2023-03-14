
package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main() {
	// creates a new server instance
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// creates a new handle function the that
	// invokes when using the "/parse" path
	http.HandleFunc("/parse", processTempl)
	fmt.Println("Server staring...")
	server.ListenAndServe()
}


func processTempl(w http.ResponseWriter, r *http.Request){
	// creates a template that parses a file named "tmpl"
	// there is no error handling
	template, _ := template.ParseFiles("t1.html", "t2.html")

	// creates the template/templatr enginie that will write
	// to the client according to the action in the file that
	// was passed.
	template.Execute(w,"Hello World!")
	}
