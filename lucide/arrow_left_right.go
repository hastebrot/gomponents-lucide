package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeftRight(children ...g.Node) g.Node {
	svg := `<path d="M8 3 4 7l4 4" /><path d="M4 7h16" /><path d="m16 21 4-4-4-4" /><path d="M20 17H4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
