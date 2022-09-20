package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse "common" templates
	tpl := template.Must(template.ParseGlob("templates/common/*.tmpl"))

	// Parse the "target" template
	tpl.ParseFiles("templates/index.tmpl")

	// You can see the template map for debugging purposes.
	fmt.Println(tpl.DefinedTemplates())

	// Execute to render to template.
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
