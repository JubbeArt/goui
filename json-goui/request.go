package main

type UiRequest struct {
	Widget *Widget  `json:"ui"`
	Events []string `json:"events"`
	Window *Window  `json:"window"`
}

type Widget struct {
	Type     *string  `json:"type"`
	Text     *string  `json:"text"`
	Image    *string  `json:"image"`
	Children []Widget `json:"children"`
}

type Window struct {
	InitialWidth  *int    `json:"initial_width"`
	InitialHeight *int    `json:"initial_height"`
	Title         *string `json:"title"`
}
