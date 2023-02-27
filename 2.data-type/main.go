package main

import "fmt"

var foo_out = "Outside main function"

func main() {
	fmt.Println("------------------------- 2. Variables and data types ------------------------- ")

	// declaring a variable: var [name] [type]

	var foo string = "Go is awesome"

	fmt.Println("\tSingle variable: ", foo)

	// multiple declartions:
	var foo2, bar string = "Hello", "World"
	// or
	var (
		foo3 string = "Hello2"
		bar2 string = "World2"
	)

	var foo4 = "What's my type?"

	// Shorthand only works inside function bodies.
	foo5 := "Shorthand!"

	fmt.Println("\tMultiple variables: ", foo2, bar)
	fmt.Println("\tMultiple variables 2: ", foo3, bar2)

	fmt.Println("\tType omitted: ", foo4)
	fmt.Println("\tShorthand variable: ", foo5)
	fmt.Println("\toutside variable: ", foo_out)

	const constant string = "This my constant"

	fmt.Println("\tMy constant: ", constant)

	var str_multiple string = `I am statically typed. 
								I was designed at Google.`
	fmt.Println("\tString multiple lines: ", str_multiple)

	var is_true bool = true
	var is_false bool = false
	fmt.Println("\tIs true: ", is_true)
	fmt.Println("\tIs false: ", is_false)

	var int_8 int8 = 127
	fmt.Println("\tint 8 byte", int_8)

	var byte_value byte = 'a'
	var rune_value rune = '🍕'

	fmt.Println("\tByte value: ", byte_value)
	fmt.Println("\tRune value: ", rune_value)

	var double_pi float32 = 3.14 * 2

	fmt.Println("\t2*Pi number: ", double_pi)
}
