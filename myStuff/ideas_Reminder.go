package main

import (
	"fmt"
	"io"
	"net/mail"
	"time"
)

func main() {

	var reminds = make([]string, 5)
	// Here, 5 is the 'length', not the capacity.
	// Adding a second int parameter would signal the capacity.

	type remModule struct {
		reminder map[struct{}]string
	}

	type date struct {
		day   time.Weekday
		month time.Month
	}

	type notice struct {
		popUp io.Writer
		eMail mail.Message
	}
	// --> Goofing Off
	reminds[0] = "Do the laundry"
	fmt.Println(reminds)
	whoa := date{day: time.Monday, month: time.April}
	fmt.Println(whoa)

	var tryIt = remModule{}
	//tryIt.reminder =
	fmt.Println(tryIt)
}
