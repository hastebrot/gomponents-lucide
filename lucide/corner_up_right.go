package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerUpRight(children ...g.Node) g.Node {
	svg := `<polyline points="15 14 20 9 15 4" /><path d="M4 20v-7a4 4 0 0 1 4-4h12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
