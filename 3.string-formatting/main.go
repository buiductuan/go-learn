// doc:https://pkg.go.dev/fmt

package main

import "fmt"

func main() {
	fmt.Println("------------------ 3. String formatting ------------------")

	name := "Golang"

	fmt.Printf("My name is %s\n", name)

	percent := (7.0 / 9) * 100

	fmt.Printf("%.2f %%", percent)
}
