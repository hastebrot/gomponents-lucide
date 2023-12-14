package lucide

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Icon(children ...g.Node) g.Node {
	return h.SVG(
		g.Group(children),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("aria-hidden", "true"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	)
}
