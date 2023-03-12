package main



import (
	"net/http"
	"html/template"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/parse", processTempl)
	server.ListenAndServe()
	
}


func processTempl(w http.ResponseWriter, r *http.Request){
	tmpl := `<!DOCTYPE html>
        <html>
          <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <title>Go Web Programming</title>
          </head>
          <body>
            {{ . }}
          </body>
        </html>`

	t := template.New("tmpl.html")
	t ,_= t.Parse(tmpl)
	t.Execute(w, "Hello Go")
}
