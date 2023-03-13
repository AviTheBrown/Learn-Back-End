
package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/parse", processTempl)
	fmt.Println("Server staring...")
	server.ListenAndServe()
}


func processTempl(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl.html")

	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)

}
