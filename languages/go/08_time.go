package main

import "fmt"
import "time"

func main() {
	// https://gobyexample.com/time-formatting-parsing
	p := fmt.Println

	t := time.Now()
	p("human readable#1", t.Format(time.RFC3339))

	p("human readable#2", t.Format(time.RFC822Z))

	loc2, _ := time.LoadLocation("America/New_York")
	p("human readable#NY", t.In(loc2).Format(time.RFC822Z))

	loc3, _ := time.LoadLocation("America/Los_Angeles")
	p("human readable#LA", t.In(loc3).Format(time.RFC822Z))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// convert nanosecond to human readable form...
	t22 := time.Unix(0, 1510610220000000000)
	p("human readable#3", t22.Format(time.RFC822Z))

	// t2 := 1508882400000000000
	// p(t2.Format(time.RFC3339))
	// p("human reaable", time.Parse(time.RFC3339, t2))

	// parse an int (timestamp) to a time type (nanosecond precision)
	t3 := int64(1508882400000000000)
	t4 := time.Unix(t3, 0)                  // assume input is ms
	t5 := time.Unix(0, t3)                  // assume input is nanoseconds
	t6 := time.Unix(0, t3).UTC()            // format it as utc as well
	t7 := time.Unix(0, t3).UTC().UnixNano() // format as utc, then back to nano (will be same) it to utc as well

	// change timezone
	//init the loc
	loc, _ := time.LoadLocation("America/New_York")
	t8 := time.Unix(0, t3).In(loc)

	// https://siongui.github.io/2017/02/16/go-parse-utime-timestamp/

	p(t4)
	p(t5)
	p(t6)
	p(t7)
	p("TZ NY", t8)
	// Predefined constants in the package implement common layouts.
	// fmt.Println("Unix format:", t.Format(time.UnixDate))
	// fmt.Println("Unix format:", t.Format(time.UnixDate))
}
