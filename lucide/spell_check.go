package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SpellCheck(children ...g.Node) g.Node {
	svg := `<path d="m6 16 6-12 6 12" /><path d="M8 12h8" /><path d="m16 20 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
