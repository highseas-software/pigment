package main

import (
	"fmt"

	"github.com/scallywaag/pigment"
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

	// ANSI 256 FG and BG
	fmt.Println(pigment.WithFG(111).String("this is an ansi foreground"))
	fmt.Println(pigment.WithBG(219).String("this is an ansi background"))

	// Formatting funcs
	pigment.WithRed().Println("red string")
	pigment.WithBold().Printf("bold string with %s\n", "formatting")
}
