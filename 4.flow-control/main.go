package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------- 4. Flow control ---------------")
	// If/else
	x := 10
	if x > 5 {
		fmt.Println("x is gt 5")
	} else if x > 10 {
		fmt.Println("x is gt 10")
	} else {
		fmt.Println("else case")
	}
	// compact if

	if y := 10; y > 5 {
		fmt.Println("y is gt 5")
	}

	//switch/case
	day := "friday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("Let's party!")
		fallthrough // get next case
	case "tuesday":
		fmt.Println("Let's party too!")
	default:
		fmt.Println("browse memes")
	}

	// compact switch

	switch c := "done"; c {
	case "done":
		fmt.Println("Work done!")
	case "failed":
		fmt.Println("Work failed!")
	default:
		fmt.Println("Something wrong!")
	}

	z := 10
	switch {
	case z > 5:
		fmt.Println("z is greater")
	default:
		fmt.Println("z is not greater")
	}

	// loop
	for i := 0; i < 10; i++ {
		// skip print
		if i < 2 {
			continue
		}
		fmt.Println("Loop i: ", i)

		// stop loop
		if i > 5 {
			break
		}
	}

	j := 0
	for j < 10 {
		fmt.Println("Loop j: ", j)
		j += 1
	}

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("infinite loop")
	}
}
