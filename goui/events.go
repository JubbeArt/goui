package goui

import (
	"fmt"
	"strconv"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type KeyEvent struct {
	Action Action
	Key    glfw.Key
	Ctrl   bool
	Shift  bool
	Alt    bool
	Super  bool
}

func (ev KeyEvent) String() string {
	keyName := glfw.GetKeyName(ev.Key, glfw.GetKeyScancode(ev.Key))
	if keyName == "" {
		keyName = "(" + strconv.Itoa(int(ev.Key)) + ")"
	}
	return fmt.Sprintf("Key: %v %v, Ctrl: %v, Shift: %v, Alt: %v, Super: %v",
		keyName, ev.Action, ev.Ctrl, ev.Shift, ev.Alt, ev.Super)
}

type ClickEvent struct {
	Button MouseButton
	X      float64
	Y      float64
	Ctrl   bool
	Shift  bool
	Alt    bool
	Super  bool
}

func (ev ClickEvent) String() string {
	return fmt.Sprintf("Click: mouse %v at pos (%.1f, %.1f), Ctrl: %v, Shift: %v, Alt: %v, Super: %v",
		ev.Button+1, ev.X, ev.Y, ev.Ctrl, ev.Shift, ev.Alt, ev.Super)
}

type MouseButton int

const (
	MouseButton1      MouseButton = MouseButton(glfw.MouseButton1)
	MouseButton2      MouseButton = MouseButton(glfw.MouseButton2)
	MouseButton3      MouseButton = MouseButton(glfw.MouseButton3)
	MouseButton4      MouseButton = MouseButton(glfw.MouseButton4)
	MouseButton5      MouseButton = MouseButton(glfw.MouseButton5)
	MouseButton6      MouseButton = MouseButton(glfw.MouseButton6)
	MouseButton7      MouseButton = MouseButton(glfw.MouseButton7)
	MouseButton8      MouseButton = MouseButton(glfw.MouseButton8)
	MouseButtonLast   MouseButton = MouseButton(glfw.MouseButtonLast)
	MouseButtonLeft   MouseButton = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  MouseButton = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle MouseButton = MouseButton(glfw.MouseButtonMiddle)
)

type ResizeEvent struct {
	Width  int
	Height int
}

func (ev ResizeEvent) String() string {
	return fmt.Sprintf("Resize: to (%v, %v)", ev.Width, ev.Height)
}

type PositionEvent struct {
	X int
	Y int
}

func (ev PositionEvent) String() string {
	return fmt.Sprintf("Position: to (%v, %v)", ev.X, ev.Y)
}

type MouseMoveEvent struct {
	X float64
	Y float64
}

func (ev MouseMoveEvent) String() string {
	return fmt.Sprintf("Mouse move: to (%.1f, %.1f)", ev.X, ev.Y)
}

type ScrollEvent struct {
	X float64
	Y float64
}

func (ev ScrollEvent) String() string {
	return fmt.Sprintf("Scroll: diff (%v, %v)", ev.X, ev.Y)
}

type Action int

const (
	Press   Action = Action(glfw.Press)
	Release Action = Action(glfw.Release)
	Repeat  Action = Action(glfw.Repeat)
)

func (a Action) String() string {
	switch a {
	case Press:
		return "pressed"
	case Release:
		return "released"
	case Repeat:
		return "repeated"
	}
	return "unknown action"
}
