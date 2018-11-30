package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Street, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:   "Classy Stays",
					Street: "121 Sanamonica Blv",
					City:   "Los Angeles",
					Zip:    "91210",
				},
				hotel{
					Name:   "Classic Stays",
					Street: "12132 Hollywood Blv",
					City:   "Los Angeles",
					Zip:    "91211",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:   "Classy Stays",
					Street: "121 Sanamonica Blv",
					City:   "Los Angeles",
					Zip:    "91210",
				},
				hotel{
					Name:   "Classic Stays",
					Street: "12132 Hollywood Blv",
					City:   "Los Angeles",
					Zip:    "91211",
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:   "Classy Stays",
					Street: "121 Sanamonica Blv",
					City:   "Los Angeles",
					Zip:    "91210",
				},
				hotel{
					Name:   "Classic Stays",
					Street: "12132 Hollywood Blv",
					City:   "Los Angeles",
					Zip:    "91211",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
