package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowRightLeft(children ...g.Node) g.Node {
	svg := `<path d="m16 3 4 4-4 4" /><path d="M20 7H4" /><path d="m8 21-4-4 4-4" /><path d="M4 17h16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
