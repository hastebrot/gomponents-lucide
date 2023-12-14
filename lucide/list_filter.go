package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListFilter(children ...g.Node) g.Node {
	svg := `<path d="M3 6h18" /><path d="M7 12h10" /><path d="M10 18h4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
