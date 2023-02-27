package main

import "fmt"

// simple function
func myFunction(text string) {
	fmt.Println("My text: ", text)
}

func addText(a, b string) string {
	return a + " " + b
}

func addTextCount(a, b string) (string, int) {
	result := a + " " + b
	return result, len(result)
}

// named returns - not clear
func sum(a, b int) (t string, s int) {
	s = a + b
	t = fmt.Sprint(s) + " sum"
	return
}

// function as value
func printFunc() {
	fn := func() {
		fmt.Println("\nInside function1")
	}
	fn()

	func() {
		fmt.Println("Executed now!")
	}()
}

func closures() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func sum_range(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func init() {
	fmt.Println(">>>>>>>>>>>>>>> Run first")
}

func init() {
	fmt.Println(">>>>>>>>>>>>>>> Run second")
}

func main() {
	fmt.Println("------------------ 5. Functions ------------------")
	myFunction("my function text")

	text := addText("Hello", "TuanBui")
	fmt.Println(text)

	text2, count := addTextCount("Test", "count")
	fmt.Printf("Text: %s, Count: %v", text2, count)

	text3, sum_value := sum(2, 3)
	fmt.Printf("\nText: %s, Sum: %v", text3, sum_value)

	printFunc()
	sum := closures()
	sum(5)
	sum(10)
	fmt.Println(sum(10))

	sum_values := sum_range(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum range result: ", sum_values)

	// defer - like stack last in first out
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("defer 3")
}
