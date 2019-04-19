package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{Name: "Tabby", Age: 21}

	tmpl := template.New("test")

	//parse some content and generate a template
	tmpl, err := tmpl.Parse("{{.Name}} is {{.Age}} years old")
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	err1 := tmpl.Execute(os.Stdout, p)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)

	}
}
//run: go run 1-accessing-structure.go
