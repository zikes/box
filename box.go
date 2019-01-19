package box

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
	"github.com/fatih/color"
	"github.com/zikes/align"
)

// Alignment determines how the contents of a Section will be aligned
type Alignment int

const (
	LeftAlign Alignment = iota
	RightAlign
	CenterAlign
)

// Box is the overall general configuration for a multi-sectionable box
type Box struct {
	// BoxStyle is the line style of the Box's outline
	*BoxStyle

	// Margin is the number of horizontal spaces before and after the Box
	Margin int

	// LineMargin is the number of vertical spaces (lines) before and after the Box
	LineMargin int

	// Padding is the number of horizontal spaces around the Box's contents
	Padding int

	// Padding is the number of vertical spaces (lines) around the Box's contents
	LinePadding int

	// BorderColor determines the color scheme used to print the Box's outline
	BorderColor *color.Color

	// ContentColor determines the color scheme used to print the Box's contents
	ContentColor *color.Color
}

// Section is a box within a Box
type Section struct {
	// BoxStyle is the line style of the Section's outline
	*BoxStyle

	// Content is the contents of the section to be printed
	Content interface{}

	// stringContent is the Content field's `string` equivalent
	stringContent string

	// BorderColor determines the color scheme used to print the Section's outline
	// If nil/omitted, it will inherit from its parent Box
	BorderColor *color.Color

	// ContentColor determines the color scheme used to print the Section's contents
	// If nil/omitted, it will inherit from its parent Box
	ContentColor *color.Color

	// Padding is the number of horizontal spaces around the Section's contents
	Padding int

	// LinePadding is the number of vertical spaces (lines) around the Section's contents
	LinePadding int

	// Alignment determines the horizontal alignment of each line of the Section's contents
	Alignment
}

// Sprint will loop through a Box's Sections and output a string of the
// Box and its overall contents.
//
// Arguments may be any mix of string and Section.
func (b *Box) Sprint(contents ...interface{}) string {
	var output strings.Builder
	sections := []Section{}

	// If no colors are specified, use no-op color operation
	if b.BorderColor == nil {
		b.BorderColor = color.New()
	}
	if b.ContentColor == nil {
		b.ContentColor = color.New()
	}

	// If no BoxStyle is specified, use DefaultStyle
	if b.BoxStyle == nil {
		b.BoxStyle = &DefaultStyle
	}

	// Determine base type of content arguments
	for _, v := range contents {
		switch t := v.(type) {
		case string:
			sections = append(sections, Section{
				BoxStyle:     b.BoxStyle,
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

	// Prepare content sections for printing
	for i, section := range sections {
		if section.BoxStyle == nil {
			sections[i].BoxStyle = b.BoxStyle
		}
		if section.BorderColor == nil {
			sections[i].BorderColor = b.BorderColor
		}
		if section.ContentColor == nil {
			sections[i].ContentColor = b.ContentColor
		}

		switch t := section.Content.(type) {
		case string:
			sections[i].stringContent = t
		case fmt.Stringer:
			sections[i].stringContent = t.String()
		}
	}

	// Determine overall internal width of box
	maxWidth := longest(sections)

	// If specified, print blank lines before printing the box
	if b.LineMargin > 0 {
		output.WriteString(strings.Repeat("\n", b.LineMargin))
	}

	// Loop through sections and stringify for output
	for idx, section := range sections {

		// Split section contents into individual line strings
		lines := strings.Split(strings.TrimSuffix(section.stringContent, "\n"), "\n")

		// If this is the first section, print the top edge of the box
		if idx == 0 {
			output.WriteString(spaces(b.Margin))
			output.WriteString(
				section.BorderColor.Sprint(
					section.BoxStyle.TopLeft,
					strings.Repeat(section.BoxStyle.Horizontal, maxWidth),
					section.BoxStyle.TopRight,
				),
			)
			output.WriteString(spaces(b.Margin) + "\n")
		}

		// Blank lines are common, define one for easy re-use
		linePad := spaces(b.Margin) +
			section.BorderColor.Sprint(section.BoxStyle.Vertical) +
			spaces(maxWidth) +
			section.BorderColor.Sprint(section.BoxStyle.Vertical) +
			spaces(b.Margin) + "\n"

		// Print a number of blank lines according to LinePadding, if any
		output.WriteString(strings.Repeat(linePad, section.LinePadding))

		for _, line := range lines {
			// Align the line contents according to the specified Section Alignment
			inner := ""
			switch section.Alignment {
			case CenterAlign:
				inner = align.Center(maxWidth-(section.Padding*2), line)
			case LeftAlign:
				inner = align.Left(maxWidth-(section.Padding*2), line)
			case RightAlign:
				inner = align.Right(maxWidth-(section.Padding*2), line)
			}
			output.WriteString(spaces(b.Margin))
			output.WriteString(section.BorderColor.Sprint(section.BoxStyle.Vertical))
			output.WriteString(
				section.ContentColor.Sprint(
					spaces(section.Padding),
					inner,
					// strlen(inner) fails to return an adequate padding because any
					// ANSI characters in the contents will count towards the length,
					// despite not being a visible printed character. The stripansi
					// library helps with this.
					spaces(maxWidth-strlen(stripansi.Strip(inner))-section.Padding*2),
					spaces(section.Padding),
				),
			)
			output.WriteString(section.BorderColor.Sprint(section.BoxStyle.Vertical))
			output.WriteString(spaces(b.Margin) + "\n")
		}

		// Print a number of blank lines according to LinePadding, if any
		output.WriteString(strings.Repeat(linePad, section.LinePadding))

		// If this is not the last section, print a horizontal box line with
		// T intersection corners
		if idx < len(sections)-1 {
			output.WriteString(spaces(b.Margin))
			output.WriteString(
				section.BorderColor.Sprint(
					section.BoxStyle.MidLeft,
					strings.Repeat(section.BoxStyle.Horizontal, maxWidth),
					section.BoxStyle.MidRight,
				),
			)
			output.WriteString(spaces(b.Margin) + "\n")
		}

		// If this is the last section, print the bottom line of the box
		if idx == len(sections)-1 {
			output.WriteString(spaces(b.Margin))
			output.WriteString(
				section.BorderColor.Sprint(
					section.BoxStyle.BottomLeft,
					strings.Repeat(section.BoxStyle.Horizontal, maxWidth),
					section.BoxStyle.BottomRight,
				),
			)
			output.WriteString(spaces(b.Margin) + "\n")
		}
	}

	// If specified, print blank lines after printing the box
	if b.LineMargin > 0 {
		output.WriteString(strings.Repeat("\n", b.LineMargin))
	}

	return output.String()
}

// Loop through sections and their content to determine the maximum width
// of the Box's contents
func longest(sections []Section) int {
	max := 0
	for _, section := range sections {
		split := strings.Split(section.stringContent, "\n")
		for _, l := range split {
			if strlen(stripansi.Strip(l))+(section.Padding*2) > max {
				max = strlen(stripansi.Strip(l)) + (section.Padding * 2)
			}
		}
	}
	return max
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// shorthand since this operation is repeated so often
func spaces(w int) string {
	if w < 0 {
		return ""
	}
	return strings.Repeat(" ", w)
}
