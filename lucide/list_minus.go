package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListMinus(children ...g.Node) g.Node {
	svg := `<path d="M11 12H3" /><path d="M16 6H3" /><path d="M16 18H3" /><path d="M21 12h-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
