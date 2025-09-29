package pigment

import "fmt"

func PrintColors() {
	fmt.Println("ANSI 256 Color Palette:")
	for i := range 256 {
		fmt.Printf("\x1b[48;5;%dm  \x1b[0m", i)
		fmt.Printf("\x1b[38;5;%dm%03d\x1b[0m ", i, i)
		if (i % 6) == 3 {
			fmt.Println()
		}
	}
}

func (c *composer) WithFG(colorID uint8) style {
	code := fmt.Sprintf("\x1b[38;5;%dm", colorID)
	return createComposer(c, code, "\x1b[39m")
}

func WithFG(colorID uint8) style {
	return defaultComposer.WithFG(colorID)
}

func (c *composer) WithBG(colorID uint8) style {
	code := fmt.Sprintf("\x1b[48;5;%dm", colorID)
	return createComposer(c, code, "\x1b[49m")
}

func WithBG(colorID uint8) style {
	return defaultComposer.WithBG(colorID)
}
