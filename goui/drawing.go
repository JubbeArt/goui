package goui

import (
	"image/color"

	"github.com/kjk/flex"
	"github.com/shibukawa/nanovgo"
)

func drawRect(ctx *nanovgo.Context, parentX, parentY float32, l *flex.Node, s Styles) {
	x := parentX + l.LayoutGetLeft()
	y := parentY + l.LayoutGetTop()
	w := l.LayoutGetWidth()
	h := l.LayoutGetHeight()
	_, _, _, _ = x, y, w, h

	background := s.background
	if background == nil {
		background = defaultBackground
	}

	ctx.BeginPath()
	ctx.SetFillColor(colorToNanoColor(background))
	ctx.RoundedRect(x, y, w, h, float32(s.borderRadius))
	ctx.Fill()
}

func drawText(ctx *nanovgo.Context, text string, parentX, parentY float32, l *flex.Node, s Styles) {
	x := parentX + l.LayoutGetLeft()
	y := parentY + l.LayoutGetTop()

	textAlign := nanovgo.Align(0)

	if s.textAlign == TextRight {
		textAlign |= nanovgo.AlignRight
		x += l.LayoutGetWidth()
	} else if s.textAlign == TextCenter {
		textAlign |= nanovgo.AlignCenter
		x += l.LayoutGetWidth() / 2
	} else {
		textAlign |= nanovgo.AlignLeft
	}

	if s.textBaseline == TextTop {
		textAlign |= nanovgo.AlignTop
	} else if s.textBaseline == TextBottom {
		textAlign |= nanovgo.AlignBottom
		y += l.LayoutGetHeight()
	} else {
		textAlign |= nanovgo.AlignMiddle
		y += l.LayoutGetHeight() / 2
	}

	col := s.color
	if col == nil {
		col = defaultColor
	}
	// TODO
	//ctx.SetTextLetterSpacing()
	//ctx.SetFontBlur()
	ctx.SetTextAlign(textAlign)
	ctx.SetFillColor(colorToNanoColor(col))
	ctx.SetFontFace(getFontFamily(s.fontFamily))
	ctx.SetFontSize(float32(getFontSize(s.fontSize)))
	ctx.Text(x, y, text)
}

func colorToNanoColor(c color.Color) nanovgo.Color {
	r, g, b, a := c.RGBA()
	const div = 256
	return nanovgo.RGBA(uint8(r/div), uint8(g/div), uint8(b/div), uint8(a/div))
}
