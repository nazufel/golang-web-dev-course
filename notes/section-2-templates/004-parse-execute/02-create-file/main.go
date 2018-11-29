package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// parse files and assign to var tpl
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// create new file name index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}

	// close nf file at end
	defer nf.Close()

	// write the template to the nf file
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
