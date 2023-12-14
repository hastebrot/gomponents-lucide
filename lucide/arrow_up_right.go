package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpRight(children ...g.Node) g.Node {
	svg := `<path d="M7 7h10v10" /><path d="M7 17 17 7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
