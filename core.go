package pigment

import "strings"

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
}

type style interface {
	String(...string) string
	WithRed() style
	Red(...string) string
	WithBold() style
	Bold(...string) string
}

var _ style = (*composer)(nil)

var defaultComposer = &composer{node: nil}

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

func (c *composer) String(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	var str string
	if len(strs) == 1 {
		str = strs[0]
	} else {
		str = strings.Join(strs, " ")
	}

	style := c.node
	if style == nil {
		return str
	}

	codeAcc := style.acc.code
	resetAcc := style.acc.reset

	if strings.Contains(str, "\x1b") {
		for style != nil {
			str = strings.ReplaceAll(str, style.style.reset, style.style.code)
			style = style.parent
		}
	}

	return codeAcc + str + resetAcc
}

func (c *composer) WithRed() style {
	return createComposer(c, "\x1b[31m", "\x1b[39m")
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
	return createComposer(c, "\x1b[1m", "\x1b[22m")
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
