package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BookKey(children ...g.Node) g.Node {
	svg := `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H14" /><path d="M20 8v14H6.5a2.5 2.5 0 0 1 0-5H20" /><circle cx="14" cy="8" r="2" /><path d="m20 2-4.5 4.5" /><path d="m19 3 1 1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
