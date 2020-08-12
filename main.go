package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"./goui"
)

func main() {
	items := []string{"Board /b/ - Random", "Board /g/ - Technology", "Board /pol/ - Politically incorrect"}
	index := 2

	err := goui.Render(func(g goui.UI) {
		g.OnClick(func(ev goui.ClickEvent) {
			fmt.Println(ev)
		})
		// and maybe mouse move

		g.Text("Jux").
			Styles(titleStyles).
			Click(func(ev goui.ClickEvent) {
				fmt.Println("button clicked", ev)
			})

		g.Box(func() {
			for i, item := range items {
				selected := i == index
				g.Box(func() {
					g.Text(item + " " + strconv.FormatBool(selected)).Styles(menuItemText)
				}).Styles(menuItem, menuItemSelected(selected))
			}
		}).Styles(menuContainer)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// or parse this stuff from css files?

var (
	textColor     = color.RGBA{R: 255, G: 255, B: 255, A: 153}
	selectedColor = color.RGBA{R: 57, G: 181, B: 74, A: 255}
)

var (
	titleStyles = goui.NewStyles().
			Color(textColor).
			FontSize(200).
			TextAlign(goui.TextCenter).
			FontFamily(goui.FontBold).
			Margin(goui.EdgeAll, 40).
			Padding(goui.EdgeAll, 20)

	menuContainer = goui.NewStyles().
			MinWidth(400).
			Height(600).
			AlignSelf(goui.AlignCenter)

	menuItem = goui.NewStyles().
			BorderRadius(5)

	menuItemSelected = goui.ConditionalStyles(
		goui.NewStyles().Background(selectedColor),
		goui.NewStyles())

	menuItemText = goui.NewStyles().
			FontSize(20).
			FontFamily(goui.FontMonoRegular).
			Color(textColor).
			Margin(goui.EdgeVertical, 6).
			Margin(goui.EdgeHorizontal, 15)
)
