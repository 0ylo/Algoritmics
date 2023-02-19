package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	amount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	perimetr := 4 * n

	// for max
	minAmount := 4 * amount
	remainder := perimetr - minAmount
	if remainder < 0 {
		fmt.Println("too many amount")
	}

	// increasing the first minimum to maximum
	max := (remainder + 4) / 4
	fmt.Println("max is: ", max*max+(amount-1))

	// for min
	var iterationCount int
	increaseCount := remainder / 4
	if increaseCount > amount {
		iterationCount = increaseCount / amount
		tenths := increaseCount % amount
		min := (iterationCount + 1) * amount
		if tenths != 0 {
			// add tens
			fmt.Println("min is :", min*amount+tenths)
		} else {
			fmt.Println("min is :", min*amount)
		}
		return
	}

	min := iterationCount*4 + (amount - iterationCount)
	fmt.Println("min is :", min*min)
}
