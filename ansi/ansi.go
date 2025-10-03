package ansi

type ansiPair struct {
	Code  uint8
	Reset uint8
}

// Modifiers
var (
	Reset         = ansiPair{0, 0}
	Bold          = ansiPair{1, 22}
	Dim           = ansiPair{2, 22}
	Italic        = ansiPair{3, 23}
	Underline     = ansiPair{4, 24}
	Inverse       = ansiPair{7, 27}
	Hidden        = ansiPair{8, 28}
	Strikethrough = ansiPair{9, 29}
)

// Colors
var (
	Black   = ansiPair{30, 39}
	Red     = ansiPair{31, 39}
	Green   = ansiPair{32, 39}
	Yellow  = ansiPair{33, 39}
	Blue    = ansiPair{34, 39}
	Magenta = ansiPair{35, 39}
	Cyan    = ansiPair{36, 39}
	White   = ansiPair{37, 39}
)

// Backgrounds
var (
	BgBlack   = ansiPair{40, 49}
	BgRed     = ansiPair{41, 49}
	BgGreen   = ansiPair{42, 49}
	BgYellow  = ansiPair{43, 49}
	BgBlue    = ansiPair{44, 49}
	BgMagenta = ansiPair{45, 49}
	BgCyan    = ansiPair{46, 49}
	BgWhite   = ansiPair{47, 49}
)

// Bright Colors
var (
	BrightBlack   = ansiPair{90, 39}
	BrightRed     = ansiPair{91, 39}
	BrightGreen   = ansiPair{92, 39}
	BrightYellow  = ansiPair{93, 39}
	BrightBlue    = ansiPair{94, 39}
	BrightMagenta = ansiPair{95, 39}
	BrightCyan    = ansiPair{96, 39}
	BrightWhite   = ansiPair{97, 39}
)

// Bright Backgrounds
var (
	BgBrightBlack   = ansiPair{100, 49}
	BgBrightRed     = ansiPair{101, 49}
	BgBrightGreen   = ansiPair{102, 49}
	BgBrightYellow  = ansiPair{103, 49}
	BgBrightBlue    = ansiPair{104, 49}
	BgBrightMagenta = ansiPair{105, 49}
	BgBrightCyan    = ansiPair{106, 49}
	BgBrightWhite   = ansiPair{107, 49}
)

var Styles = []struct {
	Name string
	Pair ansiPair
}{
	// Modifiers
	{"Reset", Reset},
	{"Bold", Bold},
	{"Dim", Dim},
	{"Italic", Italic},
	{"Underline", Underline},
	{"Inverse", Inverse},
	{"Hidden", Hidden},
	{"Strikethrough", Strikethrough},

	// Colors
	{"Black", Black},
	{"Red", Red},
	{"Green", Green},
	{"Yellow", Yellow},
	{"Blue", Blue},
	{"Magenta", Magenta},
	{"Cyan", Cyan},
	{"White", White},

	// Backgrounds
	{"BrightBlack", BrightBlack},
	{"BrightRed", BrightRed},
	{"BrightGreen", BrightGreen},
	{"BrightYellow", BrightYellow},
	{"BrightBlue", BrightBlue},
	{"BrightMagenta", BrightMagenta},
	{"BrightCyan", BrightCyan},
	{"BrightWhite", BrightWhite},

	// Bright Colors
	{"BgBlack", BgBlack},
	{"BgRed", BgRed},
	{"BgGreen", BgGreen},
	{"BgYellow", BgYellow},
	{"BgBlue", BgBlue},
	{"BgMagenta", BgMagenta},
	{"BgCyan", BgCyan},
	{"BgWhite", BgWhite},

	// Bright Backgrounds
	{"BgBrightBlack", BgBrightBlack},
	{"BgBrightRed", BgBrightRed},
	{"BgBrightGreen", BgBrightGreen},
	{"BgBrightYellow", BgBrightYellow},
	{"BgBrightBlue", BgBrightBlue},
	{"BgBrightMagenta", BgBrightMagenta},
	{"BgBrightCyan", BgBrightCyan},
	{"BgBrightWhite", BgBrightWhite},
}
