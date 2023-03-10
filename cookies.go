package main
import (
    "fmt"
"net/http"
)
func setCookie (h http.ResponseWriter, r *http.Request) {
    cookie1 := http.Cookie{
        Name: "Cookie 1",
        Value: "value 1",
    }
    cookie2 := http.Cookie{
        Name: "Cookie 2",
        Value: "value 2",
    }
    http.SetCookie(h, &cookie1)
    http.SetCookie(h, &cookie2)
    fmt.Fprintln(h, "cookie cookie2222")

}

func getCookie(w http.ResponseWriter, r *http.Request)  {
    h := r.Cookies()
    for _, cookie := range h {
        fmt.Fprintf(w, "this is the name from cookie", cookie.Name)
    }
    fmt.Fprintln(w, "cookie cookie")
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
