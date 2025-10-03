package pigment

import "fmt"

func PrintColors() {
	for i := range 256 {
		bg := WithBG(uint8(i)).String("  ")
		fg := WithFG(uint8(i)).Sprintf("%03d", i)
		fmt.Printf("%s %s ", bg, fg)

		if (i % 6) == 3 {
			fmt.Println()
		}
	}
}

func (c *composer) WithFG(colorID uint8) style {
	code := fmt.Sprintf("\x1b[38;5;%dm", colorID)
	reset := fmt.Sprintf("\x1b[%dm", 39)
	return createComposer(c, code, reset)
}

func WithFG(colorID uint8) style {
	return defaultComposer.WithFG(colorID)
}

func (c *composer) WithBG(colorID uint8) style {
	code := fmt.Sprintf("\x1b[48;5;%dm", colorID)
	reset := fmt.Sprintf("\x1b[%dm", 49)
	return createComposer(c, code, reset)
}

func WithBG(colorID uint8) style {
	return defaultComposer.WithBG(colorID)
}
