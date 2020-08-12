package goui

import (
	"fmt"

	"github.com/kjk/flex"
	"github.com/shibukawa/nanovgo"
)

type widgetContainer struct {
	widget widget
	handle *Handle
	layout *flex.Node
}

type widget interface {
	render(ctx *nanovgo.Context, parentX, parentY float32, layout *flex.Node, style Styles)
}

type boxWidget struct {
	children []*widgetContainer
}

func (g *gui) Box(children func()) *Handle {
	// create new box and add it to the current container
	newBox := &boxWidget{}
	newBoxContainer, handle := g.addWidgetExtra(newBox)

	// assume the position of the current container, and add children to that container
	old := g.currentBox
	g.currentBox = newBoxContainer
	children()

	g.currentBox = old
	return handle
}

func (w *boxWidget) render(ctx *nanovgo.Context, parentX, parentY float32, l *flex.Node, s Styles) {
	drawRect(ctx, parentX, parentY, l, s)

	x := parentX + l.LayoutGetLeft()
	y := parentY + l.LayoutGetTop()

	for _, child := range w.children {
		child.widget.render(ctx, x, y, child.layout, child.handle.styles)
	}
}

func (g *gui) Text(text string) *Handle {
	return g.addWidget(&textWidget{text: text})
}

type textWidget struct {
	text string
}

func (w *textWidget) render(ctx *nanovgo.Context, parentX, parentY float32, l *flex.Node, s Styles) {
	//drawRect(ctx, parentX, parentY, l, s)
	drawText(ctx, w.text, parentX, parentY, l, s)
}

type InputState struct {
	Text      string
	Focused   bool
	CursorPos int
	OnChange  func(text string)

	Managed bool
}

func (i InputState) String() string {
	return fmt.Sprintf("InputState: (%v)", i.Text)
}
