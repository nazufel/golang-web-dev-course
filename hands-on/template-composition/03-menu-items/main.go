package main

import (
	"log"
	"os"
	"text/template"
)

type entree struct {
	Name, Description string
	Price             float64
}

type meal struct {
	Meal    string
	Entrees []entree
}

type Meal []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// No cheating on this structure.
func main() {
	m := Meal{
		meal{
			Meal: "Breakfast",
			Entrees: []entree{
				entree{
					Name:        "Breakfast Skillet",
					Description: "Eggs, bacon, fried potatoes with toast and jam.",
					Price:       7.95,
				},
				entree{
					Name:        "Pancakes",
					Description: "Buttermilk pancakes with butter and syrup.",
					Price:       5.99,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Entrees: []entree{
				entree{
					Name:        "Sandwitch",
					Description: "Roast beef with mayo, cheese, and veggies. Served with chips.",
					Price:       9.95,
				},
				entree{
					Name:        "Salad",
					Description: "Mixed greens with fruit, nuts, and balamic vinagrette dressing.",
					Price:       9.99,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Entrees: []entree{
				entree{
					Name:        "Steak",
					Description: "Steak. Nothing more, nothing less.",
					Price:       19.95,
				},
				entree{
					Name:        "Salad",
					Description: "Mixed greens with fruit, nuts, and balamic vinagrette dressing.",
					Price:       13.99,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
