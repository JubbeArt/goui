package goui

import (
	"errors"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kjk/flex"
	"github.com/shibukawa/nanovgo"
)

type UI interface {
	//Button(text string) *Handle
	// Input() *Handle
	Text(text string) *Handle
	Box(children func()) *Handle
	Rerender()
	Title(title string)
	Size() (int, int)
	Quit()

	OnKey(func(ev KeyEvent))
	OnText(func(char rune))
	OnClick(func(ev ClickEvent))
	OnResize(func(ev ResizeEvent))
	OnPositionChange(func(ev PositionEvent))
	OnFocusChange(func(focused bool))
	OnMaximizeChange(func(maximized bool))
	OnMouseMove(func(ev MouseMoveEvent))
	OnScroll(func(ev ScrollEvent))

	//Glfw() *glfw.Window
}

// argument for interface over struct: user cant create interface themselves
type gui struct {
	window     *glfw.Window
	renderFunc func(ui UI)

	root       *widgetContainer
	currentBox *widgetContainer

	ctx         *nanovgo.Context
	queueRender chan struct{}

	keyCb       func(KeyEvent)
	textCb      func(rune)
	clickCb     func(ClickEvent)
	resizeCb    func(ResizeEvent)
	positionCb  func(PositionEvent)
	focusCb     func(bool)
	maximizeCb  func(bool)
	mouseMoveCb func(MouseMoveEvent)
	scrollCb    func(ScrollEvent)
}

func Render(render func(ui UI)) error {
	// create window
	window, err := createWindow(1200, 800, "Goui")
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	window.MakeContextCurrent()
	ctx, err := nanovgo.NewContext(nanovgo.AntiAlias /* | nanovgo.StencilStrokes | nanovgo.Debug*/)
	if err != nil {
		return err
	}
	defer ctx.Delete()

	g := &gui{
		window:      window,
		renderFunc:  render,
		queueRender: make(chan struct{}, 100),
		ctx:         ctx,
	}

	keyChannel := make(chan KeyEvent, 10)
	textChannel := make(chan rune, 10)
	clickChannel := make(chan ClickEvent, 10)
	resizeChannel := make(chan ResizeEvent, 10)
	posChannel := make(chan PositionEvent, 10)
	focusChannel := make(chan bool, 10)
	maximizedChannel := make(chan bool, 10)
	mouseMoveChannel := make(chan MouseMoveEvent, 10)
	scrollChannel := make(chan ScrollEvent, 10)

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		keyChannel <- KeyEvent{
			Action: Action(action),
			Key:    key,
			Ctrl:   mods&glfw.ModControl != 0,
			Shift:  mods&glfw.ModShift != 0,
			Alt:    mods&glfw.ModAlt != 0,
			Super:  mods&glfw.ModSuper != 0,
		}
	})

	window.SetCharCallback(func(w *glfw.Window, char rune) {
		textChannel <- char
	})

	window.SetPosCallback(func(w *glfw.Window, x int, y int) {
		posChannel <- PositionEvent{X: x, Y: y}
	})

	window.SetFocusCallback(func(w *glfw.Window, focused bool) {
		focusChannel <- focused
	})

	window.SetMaximizeCallback(func(w *glfw.Window, maximized bool) {
		maximizedChannel <- maximized
	})

	window.SetCursorPosCallback(func(w *glfw.Window, x float64, y float64) {
		mouseMoveChannel <- MouseMoveEvent{X: x, Y: y}
	})

	window.SetScrollCallback(func(w *glfw.Window, x float64, y float64) {
		scrollChannel <- ScrollEvent{X: x, Y: y}
	})

	window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		resizeChannel <- ResizeEvent{
			Width:  width,
			Height: height,
		}
		g.Rerender()
	})

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		x, y := w.GetCursorPos()
		if action == glfw.Press {
			clickChannel <- ClickEvent{
				Button: MouseButton(button),
				X:      x,
				Y:      y,
				Ctrl:   mods&glfw.ModControl != 0,
				Shift:  mods&glfw.ModShift != 0,
				Alt:    mods&glfw.ModAlt != 0,
				Super:  mods&glfw.ModSuper != 0,
			}
		} else {
			// ???
		}
	})

	fonts := []string{
		FontRegular, FontItalic, FontBold, FontBoldItalic,
		FontMonoRegular, FontMonoItalic, FontMonoBold, FontMonoBoldItalic,
	}

	for _, font := range fonts {
		path := filepath.Join("fonts", font+".ttf")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not load font: %v", err)
		}
		id := ctx.CreateFontFromMemory(font, data, 0)
		if id == -1 {
			return errors.New("could not load font")
		}
	}

	// queue initial render
	g.Rerender()

	for !window.ShouldClose() {
		shouldRender := false
	loop:
		for {
			select {
			case <-g.queueRender:
				shouldRender = true
			default:
				break loop
			}
		}

		if shouldRender {
			fbWidth, fbHeight := window.GetFramebufferSize()
			winWidth, winHeight := window.GetSize()
			pixelRatio := float32(fbWidth) / float32(winWidth)

			glBeginFrame(fbWidth, fbHeight)
			ctx.BeginFrame(winWidth, winHeight, pixelRatio)

			g.render(winWidth, winHeight)

			ctx.EndFrame()
			glEndFrame()
		}

		needRerender := false
	events:
		for {
			select {
			case ev := <-keyChannel:
				if g.keyCb != nil {
					g.keyCb(ev)
					needRerender = true
				}
			case ev := <-textChannel:
				if g.textCb != nil {
					g.textCb(ev)
					needRerender = true
				}
			case ev := <-clickChannel:
				if g.clickCb != nil {
					g.clickCb(ev)
					needRerender = true
				}
			case ev := <-resizeChannel:
				if g.resizeCb != nil {
					g.resizeCb(ev)
					needRerender = true
				}
			case ev := <-posChannel:
				if g.positionCb != nil {
					g.positionCb(ev)
					needRerender = true
				}
			case ev := <-focusChannel:
				if g.focusCb != nil {
					g.focusCb(ev)
					needRerender = true
				}
			case ev := <-maximizedChannel:
				if g.maximizeCb != nil {
					g.maximizeCb(ev)
					needRerender = true
				}
			case ev := <-mouseMoveChannel:
				if g.mouseMoveCb != nil {
					g.mouseMoveCb(ev)
					needRerender = true
				}
			case ev := <-scrollChannel:
				if g.scrollCb != nil {
					g.scrollCb(ev)
					needRerender = true
				}
			default:
				break events
			}
		}

		if needRerender {
			g.Rerender()
		}

		// TODO
		time.Sleep(10 * time.Millisecond)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	return nil
}

func (g *gui) render(width, height int) {
	// reset gui state
	g.root = &widgetContainer{
		widget: &boxWidget{},
		handle: &Handle{
			styles: NewStyles(),
		},
		layout: flex.NewNode(),
	}
	g.currentBox = g.root

	// call user provided render function and populate widget tree
	g.renderFunc(g)

	// calculate layout
	g.applyStyles(g.root)
	flex.CalculateLayout(g.root.layout, float32(width), float32(height), flex.DirectionLTR)

	// render
	g.root.widget.render(g.ctx, 0, 0, g.root.layout, g.root.handle.styles)
}

// apply styles to all flex.Node objects recursively
func (g *gui) applyStyles(w *widgetContainer) {
	applyStyles(g.ctx, w.widget, w.layout, w.handle.styles)
	if box, ok := w.widget.(*boxWidget); ok {
		for _, child := range box.children {
			g.applyStyles(child)
		}
	}
}

func (g *gui) addWidgetExtra(w widget) (*widgetContainer, *Handle) {
	widgetContainer := &widgetContainer{
		widget: w,
		layout: flex.NewNode(),
		handle: &Handle{
			styles: NewStyles(),
		},
	}
	parent := g.currentBox.widget.(*boxWidget)
	// add to parent layout
	g.currentBox.layout.InsertChild(widgetContainer.layout, len(parent.children))
	// add to parent box
	parent.children = append(parent.children, widgetContainer)
	return widgetContainer, widgetContainer.handle
}

func (g *gui) addWidget(w widget) *Handle {
	_, handle := g.addWidgetExtra(w)
	return handle
}

func (g *gui) Rerender() {
	g.queueRender <- struct{}{}
}

func (g *gui) OnKey(callback func(ev KeyEvent))                 { g.keyCb = callback }
func (g *gui) OnText(callback func(char rune))                  { g.textCb = callback }
func (g *gui) OnClick(callback func(ev ClickEvent))             { g.clickCb = callback }
func (g *gui) OnResize(callback func(ev ResizeEvent))           { g.resizeCb = callback }
func (g *gui) OnPositionChange(callback func(ev PositionEvent)) { g.positionCb = callback }
func (g *gui) OnFocusChange(callback func(focused bool))        { g.focusCb = callback }
func (g *gui) OnMaximizeChange(callback func(maximized bool))   { g.maximizeCb = callback }
func (g *gui) OnMouseMove(callback func(ev MouseMoveEvent))     { g.mouseMoveCb = callback }
func (g *gui) OnScroll(callback func(ev ScrollEvent))           { g.scrollCb = callback }

func (g *gui) Quit()              { g.window.SetShouldClose(true) }
func (g *gui) Title(title string) { g.window.SetTitle(title) }
func (g *gui) Size() (int, int)   { return g.window.GetSize() }
