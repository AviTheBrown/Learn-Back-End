package main
import (
    "fmt"
"net/http"
)
func setCookie(w http.ResponseWriter, r *http.Request){
    c1 := http.Cookie {
        Name: "cookie1",
        Value: "My cookie",
        HttpOnly: true,
    }
    c2 := http.Cookie {
        Name: "cookie2",
        Value: "second cookie value",
        HttpOnly: true,
    }
    http.SetCookie(w, &c1)
    http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
    c1 , err := r.Cookie("cookie1")
    if err != nil {
        fmt.Fprintln(w, "cookie can not be found")
    }
    cs := r.Cookies()
    fmt.Fprintln(w, c1)
    fmt.Fprintln(w, cs)
}



func main() {

    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/setCookie", setCookie)
    http.HandleFunc("/getCookie", getCookie)
    fmt.Println("server .....")
    server.ListenAndServe()
}
