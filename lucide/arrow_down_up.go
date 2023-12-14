package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownUp(children ...g.Node) g.Node {
	svg := `<path d="m3 16 4 4 4-4" /><path d="M7 20V4" /><path d="m21 8-4-4-4 4" /><path d="M17 4v16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
