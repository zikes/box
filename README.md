<h1 align="center">
<pre>
╔═════════╗
║   BOX   ║
╚═════════╝</pre>
</h1>

```go
b := &box.Box{
  BoxStyle: &box.DoubleStyle,
  Margin:  1,
  Padding: 3,
}

fmt.Print(b.Sprint("BOX"))
// ╔═════════╗
// ║   BOX   ║
// ╚═════════╝
```

* [type Box](#Box)
* [type Section](#Section)
* [type BoxStyle](#BoxStyle-2)
* [type Alignment](#Alignment)

# `Box`

```go
type Box struct {
  *BoxStyle
  Margin       int
  LineMargin   int
  Padding      int
  BorderColor  *color.Color
  ContentColor *color.Color
}

func (b *Box) Sprint(contents ...interface{}) string
```

An instance of the `Box` struct is used to determine the overall
style of the Sections it contains.

### BoxStyle

Determines the line style of the box. Available styles are `DefaultStyle`,
`DoubleStyle`, `RoundedStyle`, `ClassicStyle`, and `BlankStyle`.

You can also provide your own `BoxStyle` if you prefer, so
long as it is of the [`BoxStyle`](#BoxStyle-2) struct type.

### Margin

The number of horizontal spaces before and after the `Box`.

### LineMargin

The number of vertical spaces (lines) before and after the `Box`.

### Padding

The number of horizontal spaces around the `Box`'s contents.

### LinePadding

The number of vertical spaces (lines) around the `Box`'s contents.

### BorderColor

An optional `*color.Color` from http://github.com/fatih/color, which
determines the color scheme used to print the `Box`'s outline.

### ContentColor

An optional `*color.Color` from http://github.com/fatih/color, which
determines the color scheme used to print the `Box`'s contents.

### func (*Box) Sprint

Arguments may be any mix of `string` literal or `box.Section`. Each argument
will create a new horizontal split in the box.

# `Section`

```go
type Section struct {
  *BoxStyle
  Content      interface{}
  Padding      int
  LinePadding  int
  Alignment
  BorderColor  *color.Color
  ContentColor *color.Color
}
```

A `Section` is a horizontal split in the box, which may optionally
carry its own styles.

```
┌─────────────┐
│  Section 1  │
├─────────────┤
│  Section 2  │
├─────────────┤
│  Section 3  │
└─────────────┘
```

### BoxStyle

Useful for specifying different line styles for different sections
of the box. If omitted, the `BoxStyle` of the container `Box` is inherited.

### Content

May be a `string` literal or anything conforming to the `fmt.Stringer` interface.

### Padding

The number of horizontal spaces around the `Section`'s contents.

### LinePadding

The number of vertical spaces (lines) around the `Section`'s contents.

### Alignment

Determines the horizontal alignment of the `Section`'s contents. May be
`box.LeftAlign`, `box.RightAlign`, or `box.CenterAlign`.

### BorderColor

An optional `*color.Color` from http://github.com/fatih/color, which
determines the color scheme used to print the Section's outline. If
omitted, the `BorderColor` of the container `Box` is inherited.

### ContentColor

An optional `*color.Color` from http://github.com/fatih/color, which
determines the color scheme used to print the `Section`'s contents.

# BoxStyle

```go
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
```

Each property of a `BoxStyle` is the character to be inserted in that
position when drawing the box. For example, this is the definition for
`DefaultStyle`:

```go
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
```

Note that `MidTop` and `MidBottom` are not currently used by the `Box`
library, and may safely be omitted from custom `BoxStyle` definitions.
They are included in all built-in definitions for convenience, in case
support for vertical splits should be implemented at a later date.

Available styles:

```
┌────────────────┐ ╔═══════════════╗
│  DefaultStyle  │ ║  DoubleStyle  ║
└────────────────┘ ╚═══════════════╝
╭────────────────╮ +----------------+
│  RoundedStyle  │ |  ClassicStyle  |
╰────────────────╯ +----------------+
                
  BlankStyle†   
                
```

† `BlankStyle` uses blank spaces, which can be useful for
boxes comprised of solid background colors.

# Alignment

The built-in `Alignment` constants determine how a `Section`'s contents
are aligned. They are `LeftAlign`, `RightAlign`, and `CenterAlign`.
