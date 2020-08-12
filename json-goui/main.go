package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"sync"

	"github.com/pkg/errors"

	"../goui"
)

// test {"ui": {"type": "text", "text": "Hello World"}}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	uiChannel := make(chan UiRequest, 10)
	waitFirstRequest := sync.Mutex{} // TODO: replace this
	waitFirstRequest.Lock()

	go func() {
		for stdin.Scan() {
			line := stdin.Bytes()
			var ui UiRequest
			err := json.Unmarshal(line, &ui)
			if err != nil {
				sendError(fmt.Errorf("could not parse request: %v", err))
				continue
			}

			waitFirstRequest.Unlock()
			uiChannel <- ui
		}

		if err := stdin.Err(); err != nil {
			fmt.Println("stdin", err)
		}

	}()

	waitFirstRequest.Lock()

	var latestRequest UiRequest

	err := goui.Render(func(u goui.UI) {
	requests:
		for {
			select {
			case uiReq := <-uiChannel:
				if uiReq.Window != nil && uiReq.Window.Title != nil {
					u.Title(*uiReq.Window.Title)
				}

				latestRequest = uiReq
				u.Rerender()
			default:
				break requests
			}
		}

		if latestRequest.Widget == nil {
			return
		}

		wid := latestRequest.Widget
		if wid.Type == nil {
			sendError(errors.New("widget is missing \"type\" property"))
			return
		}

		switch *wid.Type {
		case textWidget:
			if wid.Text == nil {
				sendError(errors.New("text widget is missing \"text\" property"))
				return
			}

			u.Text(*wid.Text)
		}

		render(u, *latestRequest.Widget)

	})
	if err != nil {
		sendError(err)
		os.Exit(1)
	}
}

func sendError(err error) {
	fmt.Println(err)
}

const (
	textWidget = "text"
)

func render(u goui.UI, widget Widget) {

}

var (
	textColor     = color.RGBA{R: 255, G: 255, B: 255, A: 153}
	selectedColor = color.RGBA{R: 57, G: 181, B: 74, A: 255}
)
