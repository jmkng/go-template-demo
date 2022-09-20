package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse "common" templates
	tpl := template.Must(template.ParseGlob("templates/common/*.tmpl"))

	// Parse the "target" template. Comment one and uncomment the other to see how
	// the "content" block in base.tmpl is overriden based on which file you parse.
	tpl.ParseFiles("templates/index.tmpl")
	// tpl.ParseFiles("templates/about.tmpl")

	// You can see the template map for debugging purposes.
	fmt.Println(tpl.DefinedTemplates())

	// Execute to render to template.
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
