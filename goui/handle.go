package goui

type Handle struct {
	styles  Styles
	onClick func(ev ClickEvent)
}

func (h *Handle) Styles(styles ...Styles) *Handle {
	if len(styles) > 0 {
		h.styles = CombineStyles(styles...)
	}
	return h
}

func (h *Handle) Click(callback func(ev ClickEvent)) *Handle {
	h.onClick = callback
	return h

}
