package pigment

import "github.com/scallywaag/pigment/ansi"

type composer struct {
	node            *styleNode
	fg              *composer
	bg              *composer
	reset           *composer
	bold            *composer
	dim             *composer
	italic          *composer
	underline       *composer
	inverse         *composer
	hidden          *composer
	strikethrough   *composer
	black           *composer
	red             *composer
	green           *composer
	yellow          *composer
	blue            *composer
	magenta         *composer
	cyan            *composer
	white           *composer
	brightblack     *composer
	brightred       *composer
	brightgreen     *composer
	brightyellow    *composer
	brightblue      *composer
	brightmagenta   *composer
	brightcyan      *composer
	brightwhite     *composer
	bgblack         *composer
	bgred           *composer
	bggreen         *composer
	bgyellow        *composer
	bgblue          *composer
	bgmagenta       *composer
	bgcyan          *composer
	bgwhite         *composer
	bgbrightblack   *composer
	bgbrightred     *composer
	bgbrightgreen   *composer
	bgbrightyellow  *composer
	bgbrightblue    *composer
	bgbrightmagenta *composer
	bgbrightcyan    *composer
	bgbrightwhite   *composer
}

func (c *composer) WithReset() style {
	if c.reset == nil {
		sp := createStylePair(ansi.Reset.Code, ansi.Reset.Reset)
		c.reset = createComposer(c, sp.code, sp.reset)
	}
	return c.reset
}

func (c *composer) Reset(str ...string) string {
	return c.WithReset().String(str...)
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

func (c *composer) WithDim() style {
	if c.dim == nil {
		sp := createStylePair(ansi.Dim.Code, ansi.Dim.Reset)
		c.dim = createComposer(c, sp.code, sp.reset)
	}
	return c.dim
}

func (c *composer) Dim(str ...string) string {
	return c.WithDim().String(str...)
}

func (c *composer) WithItalic() style {
	if c.italic == nil {
		sp := createStylePair(ansi.Italic.Code, ansi.Italic.Reset)
		c.italic = createComposer(c, sp.code, sp.reset)
	}
	return c.italic
}

func (c *composer) Italic(str ...string) string {
	return c.WithItalic().String(str...)
}

func (c *composer) WithUnderline() style {
	if c.underline == nil {
		sp := createStylePair(ansi.Underline.Code, ansi.Underline.Reset)
		c.underline = createComposer(c, sp.code, sp.reset)
	}
	return c.underline
}

func (c *composer) Underline(str ...string) string {
	return c.WithUnderline().String(str...)
}

func (c *composer) WithInverse() style {
	if c.inverse == nil {
		sp := createStylePair(ansi.Inverse.Code, ansi.Inverse.Reset)
		c.inverse = createComposer(c, sp.code, sp.reset)
	}
	return c.inverse
}

func (c *composer) Inverse(str ...string) string {
	return c.WithInverse().String(str...)
}

func (c *composer) WithHidden() style {
	if c.hidden == nil {
		sp := createStylePair(ansi.Hidden.Code, ansi.Hidden.Reset)
		c.hidden = createComposer(c, sp.code, sp.reset)
	}
	return c.hidden
}

func (c *composer) Hidden(str ...string) string {
	return c.WithHidden().String(str...)
}

func (c *composer) WithStrikethrough() style {
	if c.strikethrough == nil {
		sp := createStylePair(ansi.Strikethrough.Code, ansi.Strikethrough.Reset)
		c.strikethrough = createComposer(c, sp.code, sp.reset)
	}
	return c.strikethrough
}

func (c *composer) Strikethrough(str ...string) string {
	return c.WithStrikethrough().String(str...)
}

func (c *composer) WithBlack() style {
	if c.black == nil {
		sp := createStylePair(ansi.Black.Code, ansi.Black.Reset)
		c.black = createComposer(c, sp.code, sp.reset)
	}
	return c.black
}

func (c *composer) Black(str ...string) string {
	return c.WithBlack().String(str...)
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

func (c *composer) WithGreen() style {
	if c.green == nil {
		sp := createStylePair(ansi.Green.Code, ansi.Green.Reset)
		c.green = createComposer(c, sp.code, sp.reset)
	}
	return c.green
}

func (c *composer) Green(str ...string) string {
	return c.WithGreen().String(str...)
}

func (c *composer) WithYellow() style {
	if c.yellow == nil {
		sp := createStylePair(ansi.Yellow.Code, ansi.Yellow.Reset)
		c.yellow = createComposer(c, sp.code, sp.reset)
	}
	return c.yellow
}

func (c *composer) Yellow(str ...string) string {
	return c.WithYellow().String(str...)
}

func (c *composer) WithBlue() style {
	if c.blue == nil {
		sp := createStylePair(ansi.Blue.Code, ansi.Blue.Reset)
		c.blue = createComposer(c, sp.code, sp.reset)
	}
	return c.blue
}

func (c *composer) Blue(str ...string) string {
	return c.WithBlue().String(str...)
}

func (c *composer) WithMagenta() style {
	if c.magenta == nil {
		sp := createStylePair(ansi.Magenta.Code, ansi.Magenta.Reset)
		c.magenta = createComposer(c, sp.code, sp.reset)
	}
	return c.magenta
}

func (c *composer) Magenta(str ...string) string {
	return c.WithMagenta().String(str...)
}

func (c *composer) WithCyan() style {
	if c.cyan == nil {
		sp := createStylePair(ansi.Cyan.Code, ansi.Cyan.Reset)
		c.cyan = createComposer(c, sp.code, sp.reset)
	}
	return c.cyan
}

func (c *composer) Cyan(str ...string) string {
	return c.WithCyan().String(str...)
}

func (c *composer) WithWhite() style {
	if c.white == nil {
		sp := createStylePair(ansi.White.Code, ansi.White.Reset)
		c.white = createComposer(c, sp.code, sp.reset)
	}
	return c.white
}

func (c *composer) White(str ...string) string {
	return c.WithWhite().String(str...)
}

func (c *composer) WithBrightBlack() style {
	if c.brightblack == nil {
		sp := createStylePair(ansi.BrightBlack.Code, ansi.BrightBlack.Reset)
		c.brightblack = createComposer(c, sp.code, sp.reset)
	}
	return c.brightblack
}

func (c *composer) BrightBlack(str ...string) string {
	return c.WithBrightBlack().String(str...)
}

func (c *composer) WithBrightRed() style {
	if c.brightred == nil {
		sp := createStylePair(ansi.BrightRed.Code, ansi.BrightRed.Reset)
		c.brightred = createComposer(c, sp.code, sp.reset)
	}
	return c.brightred
}

func (c *composer) BrightRed(str ...string) string {
	return c.WithBrightRed().String(str...)
}

func (c *composer) WithBrightGreen() style {
	if c.brightgreen == nil {
		sp := createStylePair(ansi.BrightGreen.Code, ansi.BrightGreen.Reset)
		c.brightgreen = createComposer(c, sp.code, sp.reset)
	}
	return c.brightgreen
}

func (c *composer) BrightGreen(str ...string) string {
	return c.WithBrightGreen().String(str...)
}

func (c *composer) WithBrightYellow() style {
	if c.brightyellow == nil {
		sp := createStylePair(ansi.BrightYellow.Code, ansi.BrightYellow.Reset)
		c.brightyellow = createComposer(c, sp.code, sp.reset)
	}
	return c.brightyellow
}

func (c *composer) BrightYellow(str ...string) string {
	return c.WithBrightYellow().String(str...)
}

func (c *composer) WithBrightBlue() style {
	if c.brightblue == nil {
		sp := createStylePair(ansi.BrightBlue.Code, ansi.BrightBlue.Reset)
		c.brightblue = createComposer(c, sp.code, sp.reset)
	}
	return c.brightblue
}

func (c *composer) BrightBlue(str ...string) string {
	return c.WithBrightBlue().String(str...)
}

func (c *composer) WithBrightMagenta() style {
	if c.brightmagenta == nil {
		sp := createStylePair(ansi.BrightMagenta.Code, ansi.BrightMagenta.Reset)
		c.brightmagenta = createComposer(c, sp.code, sp.reset)
	}
	return c.brightmagenta
}

func (c *composer) BrightMagenta(str ...string) string {
	return c.WithBrightMagenta().String(str...)
}

func (c *composer) WithBrightCyan() style {
	if c.brightcyan == nil {
		sp := createStylePair(ansi.BrightCyan.Code, ansi.BrightCyan.Reset)
		c.brightcyan = createComposer(c, sp.code, sp.reset)
	}
	return c.brightcyan
}

func (c *composer) BrightCyan(str ...string) string {
	return c.WithBrightCyan().String(str...)
}

func (c *composer) WithBrightWhite() style {
	if c.brightwhite == nil {
		sp := createStylePair(ansi.BrightWhite.Code, ansi.BrightWhite.Reset)
		c.brightwhite = createComposer(c, sp.code, sp.reset)
	}
	return c.brightwhite
}

func (c *composer) BrightWhite(str ...string) string {
	return c.WithBrightWhite().String(str...)
}

func (c *composer) WithBgBlack() style {
	if c.bgblack == nil {
		sp := createStylePair(ansi.BgBlack.Code, ansi.BgBlack.Reset)
		c.bgblack = createComposer(c, sp.code, sp.reset)
	}
	return c.bgblack
}

func (c *composer) BgBlack(str ...string) string {
	return c.WithBgBlack().String(str...)
}

func (c *composer) WithBgRed() style {
	if c.bgred == nil {
		sp := createStylePair(ansi.BgRed.Code, ansi.BgRed.Reset)
		c.bgred = createComposer(c, sp.code, sp.reset)
	}
	return c.bgred
}

func (c *composer) BgRed(str ...string) string {
	return c.WithBgRed().String(str...)
}

func (c *composer) WithBgGreen() style {
	if c.bggreen == nil {
		sp := createStylePair(ansi.BgGreen.Code, ansi.BgGreen.Reset)
		c.bggreen = createComposer(c, sp.code, sp.reset)
	}
	return c.bggreen
}

func (c *composer) BgGreen(str ...string) string {
	return c.WithBgGreen().String(str...)
}

func (c *composer) WithBgYellow() style {
	if c.bgyellow == nil {
		sp := createStylePair(ansi.BgYellow.Code, ansi.BgYellow.Reset)
		c.bgyellow = createComposer(c, sp.code, sp.reset)
	}
	return c.bgyellow
}

func (c *composer) BgYellow(str ...string) string {
	return c.WithBgYellow().String(str...)
}

func (c *composer) WithBgBlue() style {
	if c.bgblue == nil {
		sp := createStylePair(ansi.BgBlue.Code, ansi.BgBlue.Reset)
		c.bgblue = createComposer(c, sp.code, sp.reset)
	}
	return c.bgblue
}

func (c *composer) BgBlue(str ...string) string {
	return c.WithBgBlue().String(str...)
}

func (c *composer) WithBgMagenta() style {
	if c.bgmagenta == nil {
		sp := createStylePair(ansi.BgMagenta.Code, ansi.BgMagenta.Reset)
		c.bgmagenta = createComposer(c, sp.code, sp.reset)
	}
	return c.bgmagenta
}

func (c *composer) BgMagenta(str ...string) string {
	return c.WithBgMagenta().String(str...)
}

func (c *composer) WithBgCyan() style {
	if c.bgcyan == nil {
		sp := createStylePair(ansi.BgCyan.Code, ansi.BgCyan.Reset)
		c.bgcyan = createComposer(c, sp.code, sp.reset)
	}
	return c.bgcyan
}

func (c *composer) BgCyan(str ...string) string {
	return c.WithBgCyan().String(str...)
}

func (c *composer) WithBgWhite() style {
	if c.bgwhite == nil {
		sp := createStylePair(ansi.BgWhite.Code, ansi.BgWhite.Reset)
		c.bgwhite = createComposer(c, sp.code, sp.reset)
	}
	return c.bgwhite
}

func (c *composer) BgWhite(str ...string) string {
	return c.WithBgWhite().String(str...)
}

func (c *composer) WithBgBrightBlack() style {
	if c.bgbrightblack == nil {
		sp := createStylePair(ansi.BgBrightBlack.Code, ansi.BgBrightBlack.Reset)
		c.bgbrightblack = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightblack
}

func (c *composer) BgBrightBlack(str ...string) string {
	return c.WithBgBrightBlack().String(str...)
}

func (c *composer) WithBgBrightRed() style {
	if c.bgbrightred == nil {
		sp := createStylePair(ansi.BgBrightRed.Code, ansi.BgBrightRed.Reset)
		c.bgbrightred = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightred
}

func (c *composer) BgBrightRed(str ...string) string {
	return c.WithBgBrightRed().String(str...)
}

func (c *composer) WithBgBrightGreen() style {
	if c.bgbrightgreen == nil {
		sp := createStylePair(ansi.BgBrightGreen.Code, ansi.BgBrightGreen.Reset)
		c.bgbrightgreen = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightgreen
}

func (c *composer) BgBrightGreen(str ...string) string {
	return c.WithBgBrightGreen().String(str...)
}

func (c *composer) WithBgBrightYellow() style {
	if c.bgbrightyellow == nil {
		sp := createStylePair(ansi.BgBrightYellow.Code, ansi.BgBrightYellow.Reset)
		c.bgbrightyellow = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightyellow
}

func (c *composer) BgBrightYellow(str ...string) string {
	return c.WithBgBrightYellow().String(str...)
}

func (c *composer) WithBgBrightBlue() style {
	if c.bgbrightblue == nil {
		sp := createStylePair(ansi.BgBrightBlue.Code, ansi.BgBrightBlue.Reset)
		c.bgbrightblue = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightblue
}

func (c *composer) BgBrightBlue(str ...string) string {
	return c.WithBgBrightBlue().String(str...)
}

func (c *composer) WithBgBrightMagenta() style {
	if c.bgbrightmagenta == nil {
		sp := createStylePair(ansi.BgBrightMagenta.Code, ansi.BgBrightMagenta.Reset)
		c.bgbrightmagenta = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightmagenta
}

func (c *composer) BgBrightMagenta(str ...string) string {
	return c.WithBgBrightMagenta().String(str...)
}

func (c *composer) WithBgBrightCyan() style {
	if c.bgbrightcyan == nil {
		sp := createStylePair(ansi.BgBrightCyan.Code, ansi.BgBrightCyan.Reset)
		c.bgbrightcyan = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightcyan
}

func (c *composer) BgBrightCyan(str ...string) string {
	return c.WithBgBrightCyan().String(str...)
}

func (c *composer) WithBgBrightWhite() style {
	if c.bgbrightwhite == nil {
		sp := createStylePair(ansi.BgBrightWhite.Code, ansi.BgBrightWhite.Reset)
		c.bgbrightwhite = createComposer(c, sp.code, sp.reset)
	}
	return c.bgbrightwhite
}

func (c *composer) BgBrightWhite(str ...string) string {
	return c.WithBgBrightWhite().String(str...)
}
