package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerUpLeft(children ...g.Node) g.Node {
	svg := `<polyline points="9 14 4 9 9 4" /><path d="M20 20v-7a4 4 0 0 0-4-4H4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
