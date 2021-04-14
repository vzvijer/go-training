package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func processError(err error, message string, exit bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintln(os.Stderr)
		if exit {
			os.Exit(1)
		}
	}
}

func main() {
	csvFlag := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFlag)
	processError(err, fmt.Sprintf("Failed to open the CSV file %s", *csvFlag), true)
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	processError(err, fmt.Sprintf("Error while processing the CSV file %s", *csvFlag), true)
	for i, line := range lines {
		fmt.Printf("%d. %s", i+1, line)
		fmt.Println()
	}
}
