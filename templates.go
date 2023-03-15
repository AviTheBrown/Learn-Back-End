package main

import (
    "fmt"
    "html/template"
    "net/http"
    "time"
)

func formatData (t time.Time) string {
    layout := "2006-01-02"
    return t.Format(layout)
}
func process2(w http.ResponseWriter, r *http.Request) {
    // creates a key value pair map { map[stirng][function] }
    // that will be called once the template is executed.
    funcMap := template.FuncMap{ "fdate": formatData }
    // creates a new template .New("tmpl.html")
    // and adds the string function map (funcMap)
    t := template.New("tmpl.html").Funcs(funcMap)

    // parsees the template.
    t, _ = t.ParseFiles("tmpl.html")

    // executes the template engine.
//    By default, the Execute() method will use the time.Time object that was passed to it
//    as the "dot" in the template,
//    which means that it will be available to the template as the . variable.
//    When you pass time.Now() as the argument to Execute(), it sets the {{.}} variable to
//    the current time as a time.Time object.
    t.Execute(w, time.Now())

}
func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/parse", process2)
    fmt.Println("server starting ....")
    server.ListenAndServe()
}