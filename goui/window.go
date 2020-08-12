package goui

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func createWindow(w, h int, title string) (*glfw.Window, error) {
	// init Glfw
	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("error initializing GLFW: %v", err)
	}

	// the stencil size setting is required for the canvas to work
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)
	glfw.WindowHint(glfw.Samples, 8)

	// create window
	window, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating window: %v", err)
	}
	window.MakeContextCurrent()

	// init GL
	err = gl.Init()
	if err != nil {
		return nil, fmt.Errorf("error initializing GL: %v", err)
	}

	// set vsync on, enable multisample (if available)
	glfw.SwapInterval(0)
	gl.Enable(gl.MULTISAMPLE)

	return window, nil
}

func glBeginFrame(fbWidth, fbHeight int) {
	gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
}

func glEndFrame() {
	gl.Enable(gl.DEPTH_TEST)
}
