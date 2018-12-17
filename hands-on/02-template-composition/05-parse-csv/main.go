package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
	High float64
}

func main() {
	// define route to execute foo()
	http.HandleFunc("/", foo)
	// start the server
	http.ListenAndServe(":8080", nil)
}

// Foo method
func foo(w http.ResponseWriter, r *http.Request) {
	// parse the csv file and assign values to records var by passing the file to func prs()
	records := prs("table.csv")

	// parse the templates file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute the templates
	err = tpl.Execute(w, records)
	if err != nil {
		log.Fatalln(err)
	}
}

// PRS method that parses the file
func prs(filePath string) []Record {
	// open the file and check for an err
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	// close the file when done working with it
	defer src.Close()

	// open a new csv reader, read the file, and assign to rdr
	rdr := csv.NewReader(src)

	// read all lines of rdr into rows as slices and check for an error
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// make an array named records that is a slice of Record with a length of as many slices are in rows var
	records := make([]Record, 0, len(rows))

	// loop over rows and for each append to records in their perspective keys: Date, Open
	for i, row := range rows {
		// skip the first entry since it will be 0
		if i == 0 {
			continue
		}
		// parse entries of the Data column (row[0]), format it to "2006-01-02" and assign to Date
		date, _ := time.Parse("2006-01-02", row[0])
		// parse entries of the Open column (row[1]) and parse the bitsize at 64bit.
		open, _ := strconv.ParseFloat(row[1], 64)
		// parse entries of the Open column (row[1]) and parse the bitsize at 64bit.
		high, _ := strconv.ParseFloat(row[2], 64)

		// append parsed value to it's perspective stuct fields
		records = append(records, Record{
			Date: date,
			Open: open,
			High: high,
		})
	}

	// return records array back to be used by func foo()
	return records
}

// I had to cheat on this one, but I went through and added my own comments to the code so that I understood what was happening.
// I also added another column to parse through as a proof of understanding.
