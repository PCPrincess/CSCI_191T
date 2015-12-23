package presentation20

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
)
// Basic read from file
func main()  {
	f, err := os.Open("textfile.txt")
	if err != nil {
		log.Fatalln("Where's the file?")
	}
	defer f.Close()

	g, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Nothing to read")
	}

	fmt.Println(g)
	fmt.Println(string(g))
}

// Web Server Pattern
// Get data from request reader
// Get data from database reader
// Process data
// Sometimes send data to the database writer
// If request is a normal request, send data through a template
// converter to the response writer
// If request is an AJAX request, send data through a JSON
// converter to the response writer
// encoding/json package !!
// bts, err := json.Marshal(&data)
// err := json.Unmarshal(bts, &data)