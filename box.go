package box

import (
	"fmt"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/fatih/color"
	"github.com/zikes/align"
)

type Alignment int

const (
	LeftAlign Alignment = iota
	RightAlign
	CenterAlign
)

type Box struct {
	BoxStyle
	Margin       int
	Padding      int
	BorderColor  *color.Color
	ContentColor *color.Color
}

type Section struct {
	Content       interface{}
	stringContent string
	BorderColor   *color.Color
	ContentColor  *color.Color
	Padding       int
	LinePadding   int
	Alignment
}

func (b *Box) Print(contents ...interface{}) string {
	sections := []Section{}

	for _, v := range contents {
		switch t := v.(type) {
		case string:
			sections = append(sections, Section{
				Content:      t,
				BorderColor:  b.BorderColor,
				ContentColor: b.ContentColor,
				Padding:      b.Padding,
			})
		case Section:
			sections = append(sections, t)
		default:
			panic("unknown content type")
		}
	}

	for i, section := range sections {
		switch t := section.Content.(type) {
		case string:
			sections[i].stringContent = t
		case fmt.Stringer:
			sections[i].stringContent = t.String()
		}
	}

	maxWidth := longest(sections)
	output := ""

	for idx, section := range sections {
		lines := strings.Split(strings.TrimSuffix(section.stringContent, "\n"), "\n")

		if section.BorderColor == nil {
			section.BorderColor = b.BorderColor
		}
		if section.ContentColor == nil {
			section.ContentColor = b.ContentColor
		}

		if idx == 0 {
			output += strings.Repeat(" ", b.Margin) +
				section.BorderColor.Sprint(
					b.BoxStyle.TopLeft,
					strings.Repeat(b.BoxStyle.Horizontal, maxWidth),
					b.BoxStyle.TopRight,
				) +
				strings.Repeat(" ", b.Margin) + "\n"
		}

		linePad := strings.Repeat(" ", b.Margin) +
			section.BorderColor.Sprint(b.BoxStyle.Vertical) +
			strings.Repeat(" ", maxWidth) +
			section.BorderColor.Sprint(b.BoxStyle.Vertical) +
			strings.Repeat(" ", b.Margin) + "\n"

		output += strings.Repeat(linePad, section.LinePadding)

		for _, line := range lines {
			inner := ""
			switch section.Alignment {
			case CenterAlign:
				inner = align.Center(maxWidth-(section.Padding*2), line)
			case LeftAlign:
				inner = align.Left(maxWidth-(section.Padding*2), line)
			case RightAlign:
				inner = align.Right(maxWidth-(section.Padding*2), line)
			}
			output += strings.Repeat(" ", b.Margin) +
				section.BorderColor.Sprint(b.BoxStyle.Vertical) +
				strings.Repeat(" ", section.Padding) +
				section.ContentColor.Sprint(inner) +
				strings.Repeat(" ", maxWidth-len(stripansi.Strip(inner))-section.Padding*2) +
				strings.Repeat(" ", section.Padding) +
				section.BorderColor.Sprint(b.BoxStyle.Vertical) +
				strings.Repeat(" ", b.Margin) + "\n"
		}

		output += strings.Repeat(linePad, section.LinePadding)

		if idx < len(sections)-1 {
			output += strings.Repeat(" ", b.Margin) +
				section.BorderColor.Sprint(
					b.BoxStyle.MidLeft,
					strings.Repeat(b.BoxStyle.Horizontal, maxWidth),
					b.BoxStyle.MidRight,
				) +
				strings.Repeat(" ", b.Margin) + "\n"
		}
		if idx == len(sections)-1 {
			output += strings.Repeat(" ", b.Margin) +
				section.BorderColor.Sprint(b.BoxStyle.BottomLeft+
					strings.Repeat(b.BoxStyle.Horizontal, maxWidth)+
					b.BoxStyle.BottomRight) +
				strings.Repeat(" ", b.Margin) + "\n"
		}
	}

	return output
}

func longest(sections []Section) int {
	max := 0
	for _, section := range sections {
		split := strings.Split(section.stringContent, "\n")
		for _, l := range split {
			if len(stripansi.Strip(l))+(section.Padding*2) > max {
				max = len(stripansi.Strip(l)) + (section.Padding * 2)
			}
		}
	}
	return max
}
