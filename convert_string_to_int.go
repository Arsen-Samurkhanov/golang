package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines_testerday := "50"
	lines_today := "108"

	yesterday, err := strconv.Atoi(lines_testerday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)
}
