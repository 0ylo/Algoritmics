package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Not enough arguments!")
	}

	capacity, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	passengers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	busFull := passengers / capacity
	busAdd := passengers % capacity

	// caculate full bases
	for i := busFull; i > 0; i-- {
		if capacity%2 == 0 {
			// even places
			even(capacity, capacity)
		} else {
			// odd places
			odd(capacity, capacity)
		}
	}

	// if an additional bus is needed
	if busAdd != 0 {
		gone := capacity * busFull
		places := passengers - gone

		if capacity%2 == 0 {
			// even places
			even(capacity, places)
		} else {
			// odd places
			odd(capacity, places)
		}
	}

}

func odd(capacity, places int) {
	centr := capacity/2 + 1
	fmt.Println(centr)
	for i := 1; i <= places-1; i++ {
		if i%2 == 0 {
			//chetnoe i
			centr = centr + i
			fmt.Println(centr)
		} else {
			centr = centr - i
			fmt.Println(centr)
		}
	}
}

func even(capacity, places int) {
	centr := capacity / 2
	fmt.Println(centr)
	for i := 1; i <= places-1; i++ {
		if i%2 == 0 {
			//chetnoe i
			centr = centr - i
			fmt.Println(centr)
		} else {
			centr = centr + i
			fmt.Println(centr)
		}
	}
}
