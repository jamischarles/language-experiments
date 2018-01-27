// append, log some timestamps and data to a file...
package main

import "os"
import "fmt"
import "time"

func main() {
	timestamp := time.Now()

	text := "I am some output text\n"

	f, err := os.OpenFile("./10_output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	fmt.Fprintf(f, "[%s] %s", timestamp.Format(time.RFC822Z), text)

	printStats("restarts", len(text))
}

func printStats(metricName string, rows int) {
	fmt.Println("", metricName, rows)
}
