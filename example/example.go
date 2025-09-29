package main

import (
	"fmt"

	"github.com/highseas-software/pigment"
)

func main() {
	// Base colors and modifiers
	fmt.Println(pigment.Red("this is red text"))
	fmt.Println(pigment.Bold("this is bold text"))

	// Multiple strings
	fmt.Println(pigment.Red("one", "two", "three"))

	// Nesting
	fmt.Println(
		pigment.Red("this is red.. " +
			pigment.Bold("now also bold.. ") +
			"and now just red again."),
	)

	// Chaining
	fmt.Println(pigment.WithBold().Red("this is red and bold"))

	// 256 ANSI Colors
	pigment.PrintColors()
}
