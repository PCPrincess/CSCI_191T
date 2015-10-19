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
