package pigment

import (
	"fmt"

	"github.com/scallywaag/pigment/ansi"
)

type stylePair struct {
	code  string
	reset string
}

type styleNode struct {
	style  stylePair
	acc    stylePair
	parent *styleNode
}

type composer struct {
	node *styleNode
	red  *composer
	bold *composer
}

type style interface {
	String(...string) string
	Sprint(...any) string
	Sprintf(string, ...any) string
	Sprintln(...any) string
	Print(...any)
	Printf(string, ...any)
	Println(...any)
	WithRed() style
	Red(...string) string
	WithBold() style
	Bold(...string) string
	WithFG(uint8) style
	WithBG(uint8) style
}

var _ style = (*composer)(nil)

var defaultComposer = &composer{node: nil}

func codeToSeq(code uint8) string {
	return fmt.Sprintf("\x1b[%dm", code)
}

func createStylePair(code, reset uint8) stylePair {
	return stylePair{
		code:  codeToSeq(code),
		reset: codeToSeq(reset),
	}
}

func createComposer(parentComposer *composer, code string, reset string) *composer {
	var parent *styleNode
	if parentComposer.node != nil {
		parent = parentComposer.node
	}

	codeAcc := code
	resetAcc := reset
	if parent != nil {
		codeAcc = parent.acc.code + code
		resetAcc = reset + parent.acc.reset
	}

	return &composer{
		node: &styleNode{
			style:  stylePair{code: code, reset: reset},
			acc:    stylePair{code: codeAcc, reset: resetAcc},
			parent: parent,
		},
	}
}

func (c *composer) WithRed() style {
	if c.red == nil {
		sp := createStylePair(ansi.Red.Code, ansi.Red.Reset)
		c.red = createComposer(c, sp.code, sp.reset)
	}
	return c.red
}

func (c *composer) Red(str ...string) string {
	return c.WithRed().String(str...)
}

func WithRed() style {
	return defaultComposer.WithRed()
}

func Red(str ...string) string {
	return WithRed().String(str...)
}

func (c *composer) WithBold() style {
	if c.bold == nil {
		sp := createStylePair(ansi.Bold.Code, ansi.Bold.Reset)
		c.bold = createComposer(c, sp.code, sp.reset)
	}
	return c.bold
}

func (c *composer) Bold(str ...string) string {
	return c.WithBold().String(str...)
}

func WithBold() style {
	return defaultComposer.WithBold()
}

func Bold(str ...string) string {
	return WithBold().String(str...)
}
