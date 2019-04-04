package main

import (
  "fmt"
  "html/template"
  "net/http"
)

var tmpl = `<html>
<head>
    <title>Hello World!</title>
</head>
<body>
    {{ . }}
</body>
</html>
`
func handlerView(w http.ResponseWriter, r *http.Request) {
  t := template.New("hlw") //name of the template is hlw
  t, _ = t.Parse(tmpl) // parsing of template string
  t.Execute(w, "Hello World!")
}

func handlerT1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	tpName := t.Name()
	fmt.Println(tpName)
	t.Execute(w, "Asit")
}

func handlerT2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	tpName := "t2.html"
	fmt.Println(tpName)
	t.ExecuteTemplate(w, tpName, "Golang")
}

func handlerLayout(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles("layout.html", "content.html")
	tpName := t.Name()
	fmt.Println(tpName)
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/view", handlerView)
  http.HandleFunc("/t1", handlerT1)
  http.HandleFunc("/t2", handlerT2)
  http.HandleFunc("/layout", handlerLayout)
  
  server.ListenAndServe()
}
