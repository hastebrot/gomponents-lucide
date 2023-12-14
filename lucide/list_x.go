package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListX(children ...g.Node) g.Node {
	svg := `<path d="M11 12H3" /><path d="M16 6H3" /><path d="M16 18H3" /><path d="m19 10-4 4" /><path d="m15 10 4 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
