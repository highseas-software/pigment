package pigment

import (
	"fmt"
	"strings"
)

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

func (c *composer) Sprint(args ...any) string {
	return c.String(fmt.Sprint(args...))
}

func (c *composer) Sprintf(format string, args ...any) string {
	return c.String(fmt.Sprintf(format, args...))
}

func (c *composer) Sprintln(args ...any) string {
	return c.String(fmt.Sprintln(args...))
}

func (c *composer) Print(args ...any) {
	fmt.Print(c.Sprint(args...))
}

func (c *composer) Printf(format string, args ...any) {
	fmt.Print(c.Sprintf(format, args...))
}

func (c *composer) Println(args ...any) {
	fmt.Print(c.Sprintln(args...))
}
