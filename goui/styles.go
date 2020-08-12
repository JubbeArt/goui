package goui

import (
	"image/color"

	"github.com/kjk/flex"
	"github.com/shibukawa/nanovgo"
)

const unset = -1

type Styles struct {
	width     size
	height    size
	minWidth  size
	minHeight size
	maxWidth  size
	maxHeight size
	padding   edges
	margin    edges

	position Position
	overflow Overflow

	flexDirection  FlexDirection
	wrap           WrapType
	justifyContent Justify
	alignItems     Align
	alignContent   Align

	flexGrow   float64
	flexShrink float64
	alignSelf  Align

	fontFamily   string
	fontSize     float64
	textAlign    TextAlign
	textBaseline TextBaseline

	color        color.Color
	background   color.Color
	borderRadius float64
}

const defaultPadding = 0
const defaultMargin = 0
const defaultFontSize = 18
const defaultFont = FontRegular

var defaultColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var defaultBackground = color.RGBA{}

func NewStyles() Styles {
	return Styles{
		width:     size{unit: unset},
		height:    size{unit: unset},
		minWidth:  size{unit: unset},
		minHeight: size{unit: unset},
		maxWidth:  size{unit: unset},
		maxHeight: size{unit: unset},
		padding:   edges{unset, unset, unset, unset},
		margin:    edges{unset, unset, unset, unset},

		position: unset, // PositionRelative,
		overflow: unset, // OverflowVisible,

		flexDirection:  unset, //Column,
		wrap:           unset, //NoWrap,
		justifyContent: unset, //JustifyFlexStart,
		alignItems:     unset, //AlignStretch,
		alignContent:   unset, //AlignFlexStart,
		flexGrow:       unset, //0,
		flexShrink:     unset, //1,
		alignSelf:      unset, //AlignAuto,

		fontFamily:   "",    //"Righteous-Regular.ttf",
		fontSize:     unset, //18,
		textAlign:    unset, //TextLeft,
		textBaseline: unset, //TextMiddle,

		color:        nil, // color.RGBA{R: 255, G: 255, B: 255, A: 255},
		background:   nil, // color.RGBA{},
		borderRadius: unset,
	}
}

func getFontSize(size float64) float64 {
	if size == unset {
		return defaultFontSize
	}
	return size
}

func getFontFamily(path string) string {
	if path == "" {
		return defaultFont
	}
	return path
}

func ConditionalStyles(trueStyles, falseStyles Styles) func(cond bool) Styles {
	return func(cond bool) Styles {
		if cond {
			return trueStyles
		} else {
			return falseStyles
		}
	}
}

func (h Styles) Width(px float64) Styles              { h.width = size{px, unitPx}; return h }
func (h Styles) WidthPct(pct float64) Styles          { h.width = size{pct, unitPercent}; return h }
func (h Styles) Height(px float64) Styles             { h.height = size{px, unitPx}; return h }
func (h Styles) HeightPct(pct float64) Styles         { h.height = size{pct, unitPercent}; return h }
func (h Styles) MinWidth(px float64) Styles           { h.minWidth = size{px, unitPx}; return h }
func (h Styles) MinWidthPct(pct float64) Styles       { h.minWidth = size{pct, unitPercent}; return h }
func (h Styles) MinHeight(px float64) Styles          { h.minHeight = size{px, unitPx}; return h }
func (h Styles) MinHeightPct(pct float64) Styles      { h.minHeight = size{pct, unitPercent}; return h }
func (h Styles) MaxWidth(px float64) Styles           { h.maxWidth = size{value: px, unit: unitPx}; return h }
func (h Styles) MaxWidthPct(pct float64) Styles       { h.maxWidth = size{pct, unitPercent}; return h }
func (h Styles) MaxHeight(px float64) Styles          { h.maxHeight = size{px, unitPx}; return h }
func (h Styles) MaxHeightPct(pct float64) Styles      { h.maxHeight = size{pct, unitPercent}; return h }
func (h Styles) Margin(edge Edge, px float64) Styles  { h.margin = h.margin.apply(edge, px); return h }
func (h Styles) Padding(edge Edge, px float64) Styles { h.padding = h.padding.apply(edge, px); return h }

func (h Styles) Position(pos Position) Styles      { h.position = pos; return h }
func (h Styles) Overflow(overflow Overflow) Styles { h.overflow = overflow; return h }

func (h Styles) FlexDirection(direction FlexDirection) Styles { h.flexDirection = direction; return h }
func (h Styles) Wrap(wrap WrapType) Styles                    { h.wrap = wrap; return h }
func (h Styles) JustifyContent(justify Justify) Styles        { h.justifyContent = justify; return h }
func (h Styles) AlignItems(align Align) Styles                { h.alignItems = align; return h }
func (h Styles) AlignContent(align Align) Styles              { h.alignContent = align; return h }
func (h Styles) FlexGrow(grow float64) Styles                 { h.flexGrow = grow; return h }
func (h Styles) FlexShrink(shrink float64) Styles             { h.flexShrink = shrink; return h }
func (h Styles) AlignSelf(align Align) Styles                 { h.alignSelf = align; return h }

func (h Styles) FontFamily(path string) Styles             { h.fontFamily = path; return h }
func (h Styles) FontSize(px float64) Styles                { h.fontSize = px; return h }
func (h Styles) TextAlign(align TextAlign) Styles          { h.textAlign = align; return h }
func (h Styles) TextBaseline(baseline TextBaseline) Styles { h.textBaseline = baseline; return h }

func (h Styles) Color(color color.Color) Styles      { h.color = color; return h }
func (h Styles) Background(color color.Color) Styles { h.background = color; return h }
func (h Styles) BorderRadius(px float64) Styles      { h.borderRadius = px; return h }

func applyStyles(ctx *nanovgo.Context, widget widget, l *flex.Node, s Styles) {
	if wid, ok := widget.(*textWidget); ok {
		//ctx.cv.SetFont(getFontFamily(s.fontFamily), getFontSize(s.fontSize))
		//textWidth := ctx.cv.MeasureText(wid.text).Width
		//l.StyleSetMinWidth(float32(textWidth))
		l.StyleSetMinHeight(float32(getFontSize(s.fontSize)))
		_ = wid
	}

	setLayoutSize(s.width, l.StyleSetWidth, l.StyleSetWidthPercent)
	setLayoutSize(s.height, l.StyleSetHeight, l.StyleSetHeightPercent)
	setLayoutSize(s.minWidth, l.StyleSetMinWidth, l.StyleSetMinWidthPercent)
	setLayoutSize(s.minHeight, l.StyleSetMinHeight, l.StyleSetMinHeightPercent)
	setLayoutSize(s.maxWidth, l.StyleSetMaxWidth, l.StyleSetMaxWidthPercent)
	setLayoutSize(s.maxHeight, l.StyleSetMaxHeight, l.StyleSetMaxHeightPercent)

	setLayoutEdges(s.margin, defaultMargin, l.StyleSetMargin)
	setLayoutEdges(s.padding, defaultPadding, l.StyleSetPadding)

	l.StyleSetPositionType(flex.PositionType(checkUnset(int(s.position), int(flex.PositionTypeRelative))))
	l.StyleSetOverflow(flex.Overflow(checkUnset(int(s.overflow), int(flex.OverflowVisible))))
	l.StyleSetFlexDirection(flex.FlexDirection(checkUnset(int(s.flexDirection), int(flex.FlexDirectionColumn))))
	l.StyleSetFlexWrap(flex.Wrap(checkUnset(int(s.wrap), int(flex.WrapNoWrap))))
	l.StyleSetJustifyContent(flex.Justify(checkUnset(int(s.justifyContent), int(flex.JustifyFlexStart))))
	l.StyleSetAlignItems(flex.Align(checkUnset(int(s.alignItems), int(flex.AlignStretch))))
	l.StyleSetAlignContent(flex.Align(checkUnset(int(s.alignContent), int(flex.AlignFlexStart))))
	l.StyleSetFlexGrow(checkUnsetF(s.flexGrow, 0))
	l.StyleSetFlexShrink(checkUnsetF(s.flexShrink, 1))
	l.StyleSetAlignSelf(flex.Align(checkUnset(int(s.alignSelf), int(flex.AlignAuto))))
}

func setLayoutSize(size size, setPx, setPercent func(val float32)) {
	switch size.unit {
	case unitPx:
		setPx(float32(size.value))
	case unitPercent:
		setPercent(float32(size.value))
	}
}

func setLayoutEdges(edges edges, defaultVal float64, setEdge func(edge flex.Edge, val float32)) {
	sides := []flex.Edge{flex.EdgeTop, flex.EdgeRight, flex.EdgeBottom, flex.EdgeLeft}
	values := []float64{edges.top, edges.right, edges.bottom, edges.left}

	for i, val := range values {
		setEdge(sides[i], checkUnsetF(val, defaultVal))
	}
}

func checkUnset(val, defaultVal int) int {
	if val == unset {
		return defaultVal
	}
	return val
}

func checkUnsetF(val, defaultVal float64) float32 {
	if val == unset {
		return float32(defaultVal)
	}
	return float32(val)
}

func CombineStyles(styles ...Styles) Styles {
	if len(styles) == 0 {
		panic("goui.CombineStyles needs at least 1 argument")
	}

	style := styles[0]

	for _, s := range styles[1:] {
		if s.width.unit != unset {
			style.width = s.width
		}
		if s.height.unit != unset {
			style.height = s.height
		}

		if s.minWidth.unit != unset {
			style.minWidth = s.minWidth
		}
		if s.minHeight.unit != unset {
			style.minHeight = s.minHeight
		}

		if s.maxWidth.unit != unset {
			style.maxWidth = s.maxWidth
		}
		if s.maxHeight.unit != unset {
			style.maxHeight = s.maxHeight
		}

		if s.padding.top != unset {
			style.padding.top = s.padding.top
		}
		if s.padding.right != unset {
			style.padding.right = s.padding.right
		}
		if s.padding.bottom != unset {
			style.padding.bottom = s.padding.bottom
		}
		if s.padding.left != unset {
			style.padding.left = s.padding.left
		}

		if s.margin.top != unset {
			style.margin.top = s.margin.top
		}
		if s.margin.right != unset {
			style.margin.right = s.margin.right
		}
		if s.margin.bottom != unset {
			style.margin.bottom = s.margin.bottom
		}
		if s.margin.left != unset {
			style.margin.left = s.margin.left
		}

		if s.position != unset {
			style.position = s.position
		}
		if s.overflow != unset {
			style.overflow = s.overflow
		}

		if s.flexDirection != unset {
			style.flexDirection = s.flexDirection
		}
		if s.wrap != unset {
			style.wrap = s.wrap
		}
		if s.justifyContent != unset {
			style.justifyContent = s.justifyContent
		}
		if s.alignItems != unset {
			style.alignItems = s.alignItems
		}
		if s.alignContent != unset {
			style.alignContent = s.alignContent
		}
		if s.flexGrow != unset {
			style.flexGrow = s.flexGrow
		}
		if s.flexShrink != unset {
			style.flexShrink = s.flexShrink
		}
		if s.alignSelf != unset {
			style.alignSelf = s.alignSelf
		}

		if s.fontFamily != "" {
			style.fontFamily = s.fontFamily
		}
		if s.fontSize != unset {
			style.fontSize = s.fontSize
		}
		if s.textAlign != unset {
			style.textAlign = s.textAlign
		}
		if s.textBaseline != unset {
			style.textBaseline = s.textBaseline
		}

		if s.color != nil {
			style.color = s.color
		}
		if s.background != nil {
			style.background = s.background
		}
		if s.borderRadius != unset {
			style.borderRadius = s.borderRadius
		}
	}

	return style
}

type Edge int

const (
	EdgeAll Edge = iota
	EdgeVertical
	EdgeHorizontal
	EdgeTop
	EdgeRight
	EdgeBottom
	EdgeLeft
)

type size struct {
	value float64
	unit  unit
}

const (
	FontRegular    = "Roboto-Regular"
	FontItalic     = "Roboto-Italic"
	FontBold       = "Roboto-Bold"
	FontBoldItalic = "Roboto-BoldItalic"

	FontMonoRegular    = "RobotoMono-Regular"
	FontMonoItalic     = "RobotoMono-Italic"
	FontMonoBold       = "RobotoMono-Bold"
	FontMonoBoldItalic = "RobotoMono-BoldItalic"
)

type unit int

const (
	unitPx unit = iota
	unitPercent
)

type edges struct {
	top, right, bottom, left float64
}

func (e edges) apply(edge Edge, px float64) edges {
	switch edge {
	case EdgeAll:
		e.top = px
		e.right = px
		e.bottom = px
		e.left = px
	case EdgeVertical:
		e.top = px
		e.bottom = px
	case EdgeHorizontal:
		e.left = px
		e.right = px
	case EdgeTop:
		e.top = px
	case EdgeRight:
		e.right = px
	case EdgeBottom:
		e.bottom = px
	case EdgeLeft:
		e.left = px
	}
	return e
}

type TextAlign int

const (
	TextLeft TextAlign = iota
	TextCenter
	TextRight
)

type TextBaseline int

const (
	TextTop TextBaseline = iota
	TextMiddle
	TextBottom
)

// Copied from https://github.com/kjk/flex/blob/ed34d6b6a425cc6c1b76e224b0c882d608c12aaa/enums.go

type Align int

const (
	AlignAuto         Align = Align(flex.AlignAuto)
	AlignFlexStart    Align = Align(flex.AlignFlexStart)
	AlignCenter       Align = Align(flex.AlignCenter)
	AlignFlexEnd      Align = Align(flex.AlignFlexEnd)
	AlignStretch      Align = Align(flex.AlignStretch)
	AlignBaseline     Align = Align(flex.AlignBaseline)
	AlignSpaceBetween Align = Align(flex.AlignSpaceBetween)
	AlignSpaceAround  Align = Align(flex.AlignSpaceAround)
)

type FlexDirection int

const (
	Column        FlexDirection = FlexDirection(flex.FlexDirectionColumn)
	ColumnReverse FlexDirection = FlexDirection(flex.FlexDirectionColumnReverse)
	Row           FlexDirection = FlexDirection(flex.FlexDirectionRow)
	RowReverse    FlexDirection = FlexDirection(flex.FlexDirectionRowReverse)
)

type Justify int

const (
	JustifyFlexStart    Justify = Justify(flex.JustifyFlexStart)
	JustifyCenter       Justify = Justify(flex.JustifyCenter)
	JustifyFlexEnd      Justify = Justify(flex.JustifyFlexEnd)
	JustifySpaceBetween Justify = Justify(flex.JustifySpaceBetween)
	JustifySpaceAround  Justify = Justify(flex.JustifySpaceAround)
)

type Overflow int

const (
	OverflowVisible Overflow = Overflow(flex.OverflowVisible)
	OverflowHidden  Overflow = Overflow(flex.OverflowHidden)
	OverflowScroll  Overflow = Overflow(flex.OverflowScroll)
)

type Position int

const (
	PositionRelative Position = Position(flex.PositionTypeRelative)
	PositionAbsolute Position = Position(flex.PositionTypeAbsolute)
)

type WrapType int

const (
	NoWrap      WrapType = WrapType(flex.WrapNoWrap)
	Wrap        WrapType = WrapType(flex.WrapWrap)
	WrapReverse WrapType = WrapType(flex.WrapWrapReverse)
)

func (value TextAlign) String() string {
	switch value {
	case unset:
		return "unset"
	case TextLeft:
		return "left"
	case TextCenter:
		return "center"
	case TextRight:
		return "right"
	}
	return "unknown"
}

func (value TextBaseline) String() string {
	switch value {
	case unset:
		return "unset"
	case TextTop:
		return "top"
	case TextMiddle:
		return "middle"
	case TextBottom:
		return "bottom"
	}
	return "unknown"
}

func (value Align) String() string {
	switch value {
	case unset:
		return "unset"
	case AlignAuto:
		return "auto"
	case AlignFlexStart:
		return "flex-start"
	case AlignCenter:
		return "center"
	case AlignFlexEnd:
		return "flex-end"
	case AlignStretch:
		return "stretch"
	case AlignBaseline:
		return "baseline"
	case AlignSpaceBetween:
		return "space-between"
	case AlignSpaceAround:
		return "space-around"
	}
	return "unknown"
}

func (value FlexDirection) String() string {
	switch value {
	case unset:
		return "unset"
	case Column:
		return "column"
	case ColumnReverse:
		return "column-reverse"
	case Row:
		return "row"
	case RowReverse:
		return "row-reverse"
	}
	return "unknown"
}

func (value Justify) String() string {
	switch value {
	case unset:
		return "unset"
	case JustifyFlexStart:
		return "flex-start"
	case JustifyCenter:
		return "center"
	case JustifyFlexEnd:
		return "flex-end"
	case JustifySpaceBetween:
		return "space-between"
	case JustifySpaceAround:
		return "space-around"
	}
	return "unknown"
}

func (value Overflow) String() string {
	switch value {
	case unset:
		return "unset"
	case OverflowVisible:
		return "visible"
	case OverflowHidden:
		return "hidden"
	case OverflowScroll:
		return "scroll"
	}
	return "unknown"
}

func (value Position) String() string {
	switch value {
	case unset:
		return "unset"
	case PositionRelative:
		return "relative"
	case PositionAbsolute:
		return "absolute"
	}
	return "unknown"
}

func (value WrapType) String() string {
	switch value {
	case unset:
		return "unset"
	case NoWrap:
		return "no-wrap"
	case Wrap:
		return "wrap"
	case WrapReverse:
		return "wrap-reverse"
	}
	return "unknown"
}
