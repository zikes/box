package box

type BoxStyle struct {
	TopLeft     string
	TopRight    string
	BottomRight string
	BottomLeft  string
	MidRight    string
	MidLeft     string
	MidTop      string
	MidBottom   string
	Vertical    string
	Horizontal  string
}

var DefaultStyle BoxStyle = BoxStyle{
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomRight: "┘",
	BottomLeft:  "└",
	MidRight:    "┤",
	MidLeft:     "├",
	MidTop:      "┬",
	MidBottom:   "┴",
	Vertical:    "│",
	Horizontal:  "─",
}

var DoubleStyle BoxStyle = BoxStyle{
	TopLeft:     "╔",
	TopRight:    "╗",
	BottomRight: "╝",
	BottomLeft:  "╚",
	MidRight:    "╣",
	MidLeft:     "╠",
	MidTop:      "╦",
	MidBottom:   "╩",
	Vertical:    "║",
	Horizontal:  "═",
}

var RoundedStyle BoxStyle = BoxStyle{
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomRight: "╯",
	BottomLeft:  "╰",
	MidRight:    "┤",
	MidLeft:     "├",
	MidTop:      "┬",
	MidBottom:   "┴",
	Vertical:    "│",
	Horizontal:  "─",
}

var ClassicStyle BoxStyle = BoxStyle{
	TopLeft:     "+",
	TopRight:    "+",
	BottomRight: "+",
	BottomLeft:  "+",
	MidRight:    "+",
	MidLeft:     "+",
	MidTop:      "+",
	MidBottom:   "+",
	Vertical:    "|",
	Horizontal:  "-",
}

var BlankStyle BoxStyle = BoxStyle{
	TopLeft:     " ",
	TopRight:    " ",
	BottomRight: " ",
	BottomLeft:  " ",
	MidRight:    " ",
	MidLeft:     " ",
	MidTop:      " ",
	MidBottom:   " ",
	Vertical:    " ",
	Horizontal:  " ",
}
