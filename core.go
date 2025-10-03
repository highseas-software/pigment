package pigment

import (
	"fmt"
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
