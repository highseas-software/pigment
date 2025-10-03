package pigment

type style interface {
	String(...string) string
	Sprint(...any) string
	Sprintf(string, ...any) string
	Sprintln(...any) string
	Print(...any)
	Printf(string, ...any)
	Println(...any)
	WithFG(uint8) style
	WithBG(uint8) style
	WithReset() style
	Reset(...string) string
	WithBold() style
	Bold(...string) string
	WithDim() style
	Dim(...string) string
	WithItalic() style
	Italic(...string) string
	WithUnderline() style
	Underline(...string) string
	WithInverse() style
	Inverse(...string) string
	WithHidden() style
	Hidden(...string) string
	WithStrikethrough() style
	Strikethrough(...string) string
	WithBlack() style
	Black(...string) string
	WithRed() style
	Red(...string) string
	WithGreen() style
	Green(...string) string
	WithYellow() style
	Yellow(...string) string
	WithBlue() style
	Blue(...string) string
	WithMagenta() style
	Magenta(...string) string
	WithCyan() style
	Cyan(...string) string
	WithWhite() style
	White(...string) string
	WithBrightBlack() style
	BrightBlack(...string) string
	WithBrightRed() style
	BrightRed(...string) string
	WithBrightGreen() style
	BrightGreen(...string) string
	WithBrightYellow() style
	BrightYellow(...string) string
	WithBrightBlue() style
	BrightBlue(...string) string
	WithBrightMagenta() style
	BrightMagenta(...string) string
	WithBrightCyan() style
	BrightCyan(...string) string
	WithBrightWhite() style
	BrightWhite(...string) string
	WithBgBlack() style
	BgBlack(...string) string
	WithBgRed() style
	BgRed(...string) string
	WithBgGreen() style
	BgGreen(...string) string
	WithBgYellow() style
	BgYellow(...string) string
	WithBgBlue() style
	BgBlue(...string) string
	WithBgMagenta() style
	BgMagenta(...string) string
	WithBgCyan() style
	BgCyan(...string) string
	WithBgWhite() style
	BgWhite(...string) string
	WithBgBrightBlack() style
	BgBrightBlack(...string) string
	WithBgBrightRed() style
	BgBrightRed(...string) string
	WithBgBrightGreen() style
	BgBrightGreen(...string) string
	WithBgBrightYellow() style
	BgBrightYellow(...string) string
	WithBgBrightBlue() style
	BgBrightBlue(...string) string
	WithBgBrightMagenta() style
	BgBrightMagenta(...string) string
	WithBgBrightCyan() style
	BgBrightCyan(...string) string
	WithBgBrightWhite() style
	BgBrightWhite(...string) string
}

func WithReset() style {
	return defaultComposer.WithReset()
}

func Reset(str ...string) string {
	return WithReset().String(str...)
}

func WithBold() style {
	return defaultComposer.WithBold()
}

func Bold(str ...string) string {
	return WithBold().String(str...)
}

func WithDim() style {
	return defaultComposer.WithDim()
}

func Dim(str ...string) string {
	return WithDim().String(str...)
}

func WithItalic() style {
	return defaultComposer.WithItalic()
}

func Italic(str ...string) string {
	return WithItalic().String(str...)
}

func WithUnderline() style {
	return defaultComposer.WithUnderline()
}

func Underline(str ...string) string {
	return WithUnderline().String(str...)
}

func WithInverse() style {
	return defaultComposer.WithInverse()
}

func Inverse(str ...string) string {
	return WithInverse().String(str...)
}

func WithHidden() style {
	return defaultComposer.WithHidden()
}

func Hidden(str ...string) string {
	return WithHidden().String(str...)
}

func WithStrikethrough() style {
	return defaultComposer.WithStrikethrough()
}

func Strikethrough(str ...string) string {
	return WithStrikethrough().String(str...)
}

func WithBlack() style {
	return defaultComposer.WithBlack()
}

func Black(str ...string) string {
	return WithBlack().String(str...)
}

func WithRed() style {
	return defaultComposer.WithRed()
}

func Red(str ...string) string {
	return WithRed().String(str...)
}

func WithGreen() style {
	return defaultComposer.WithGreen()
}

func Green(str ...string) string {
	return WithGreen().String(str...)
}

func WithYellow() style {
	return defaultComposer.WithYellow()
}

func Yellow(str ...string) string {
	return WithYellow().String(str...)
}

func WithBlue() style {
	return defaultComposer.WithBlue()
}

func Blue(str ...string) string {
	return WithBlue().String(str...)
}

func WithMagenta() style {
	return defaultComposer.WithMagenta()
}

func Magenta(str ...string) string {
	return WithMagenta().String(str...)
}

func WithCyan() style {
	return defaultComposer.WithCyan()
}

func Cyan(str ...string) string {
	return WithCyan().String(str...)
}

func WithWhite() style {
	return defaultComposer.WithWhite()
}

func White(str ...string) string {
	return WithWhite().String(str...)
}

func WithBrightBlack() style {
	return defaultComposer.WithBrightBlack()
}

func BrightBlack(str ...string) string {
	return WithBrightBlack().String(str...)
}

func WithBrightRed() style {
	return defaultComposer.WithBrightRed()
}

func BrightRed(str ...string) string {
	return WithBrightRed().String(str...)
}

func WithBrightGreen() style {
	return defaultComposer.WithBrightGreen()
}

func BrightGreen(str ...string) string {
	return WithBrightGreen().String(str...)
}

func WithBrightYellow() style {
	return defaultComposer.WithBrightYellow()
}

func BrightYellow(str ...string) string {
	return WithBrightYellow().String(str...)
}

func WithBrightBlue() style {
	return defaultComposer.WithBrightBlue()
}

func BrightBlue(str ...string) string {
	return WithBrightBlue().String(str...)
}

func WithBrightMagenta() style {
	return defaultComposer.WithBrightMagenta()
}

func BrightMagenta(str ...string) string {
	return WithBrightMagenta().String(str...)
}

func WithBrightCyan() style {
	return defaultComposer.WithBrightCyan()
}

func BrightCyan(str ...string) string {
	return WithBrightCyan().String(str...)
}

func WithBrightWhite() style {
	return defaultComposer.WithBrightWhite()
}

func BrightWhite(str ...string) string {
	return WithBrightWhite().String(str...)
}

func WithBgBlack() style {
	return defaultComposer.WithBgBlack()
}

func BgBlack(str ...string) string {
	return WithBgBlack().String(str...)
}

func WithBgRed() style {
	return defaultComposer.WithBgRed()
}

func BgRed(str ...string) string {
	return WithBgRed().String(str...)
}

func WithBgGreen() style {
	return defaultComposer.WithBgGreen()
}

func BgGreen(str ...string) string {
	return WithBgGreen().String(str...)
}

func WithBgYellow() style {
	return defaultComposer.WithBgYellow()
}

func BgYellow(str ...string) string {
	return WithBgYellow().String(str...)
}

func WithBgBlue() style {
	return defaultComposer.WithBgBlue()
}

func BgBlue(str ...string) string {
	return WithBgBlue().String(str...)
}

func WithBgMagenta() style {
	return defaultComposer.WithBgMagenta()
}

func BgMagenta(str ...string) string {
	return WithBgMagenta().String(str...)
}

func WithBgCyan() style {
	return defaultComposer.WithBgCyan()
}

func BgCyan(str ...string) string {
	return WithBgCyan().String(str...)
}

func WithBgWhite() style {
	return defaultComposer.WithBgWhite()
}

func BgWhite(str ...string) string {
	return WithBgWhite().String(str...)
}

func WithBgBrightBlack() style {
	return defaultComposer.WithBgBrightBlack()
}

func BgBrightBlack(str ...string) string {
	return WithBgBrightBlack().String(str...)
}

func WithBgBrightRed() style {
	return defaultComposer.WithBgBrightRed()
}

func BgBrightRed(str ...string) string {
	return WithBgBrightRed().String(str...)
}

func WithBgBrightGreen() style {
	return defaultComposer.WithBgBrightGreen()
}

func BgBrightGreen(str ...string) string {
	return WithBgBrightGreen().String(str...)
}

func WithBgBrightYellow() style {
	return defaultComposer.WithBgBrightYellow()
}

func BgBrightYellow(str ...string) string {
	return WithBgBrightYellow().String(str...)
}

func WithBgBrightBlue() style {
	return defaultComposer.WithBgBrightBlue()
}

func BgBrightBlue(str ...string) string {
	return WithBgBrightBlue().String(str...)
}

func WithBgBrightMagenta() style {
	return defaultComposer.WithBgBrightMagenta()
}

func BgBrightMagenta(str ...string) string {
	return WithBgBrightMagenta().String(str...)
}

func WithBgBrightCyan() style {
	return defaultComposer.WithBgBrightCyan()
}

func BgBrightCyan(str ...string) string {
	return WithBgBrightCyan().String(str...)
}

func WithBgBrightWhite() style {
	return defaultComposer.WithBgBrightWhite()
}

func BgBrightWhite(str ...string) string {
	return WithBgBrightWhite().String(str...)
}
