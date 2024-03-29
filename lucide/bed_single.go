package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BedSingle(children ...g.Node) g.Node {
	svg := `<path d="M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8" /><path d="M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4" /><path d="M3 18h18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
