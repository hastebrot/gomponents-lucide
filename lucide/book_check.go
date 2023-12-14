package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BookCheck(children ...g.Node) g.Node {
	svg := `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20" /><path d="m9 9.5 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
