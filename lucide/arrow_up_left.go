package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	svg := `<path d="M7 17V7h10" /><path d="M17 17 7 7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
