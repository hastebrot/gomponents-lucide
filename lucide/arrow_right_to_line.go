package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowRightToLine(children ...g.Node) g.Node {
	svg := `<path d="M17 12H3" /><path d="m11 18 6-6-6-6" /><path d="M21 5v14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
