package main

import (
	"fmt"
	"io"
	"net/mail"
	"time"
)

func main()  {

	var reminds = make([]string, 5)
	// Need to re-read slices, obviously

	type remModule struct {
		reminder map[struct {}]string
	}

	type date struct {
		day time.Weekday
		month time.Month
		year time.Time
	}

	type notice struct {
		popUp io.Writer
		eMail mail.Message
	}
// --> Goofing Off
	reminds[0] = "Do the laundry"
	fmt.Println(reminds)
	whoa := date{day:time.Monday, month:time.April, year:time.Now() }
	fmt.Println(whoa)
}
