package main

import (
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

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/view", handlerView)
  server.ListenAndServe()
}
