package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Repeat1(children ...g.Node) g.Node {
	svg := `<path d="m17 2 4 4-4 4" /><path d="M3 11v-1a4 4 0 0 1 4-4h14" /><path d="m7 22-4-4 4-4" /><path d="M21 13v1a4 4 0 0 1-4 4H3" /><path d="M11 10h1v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
