package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerDownLeft(children ...g.Node) g.Node {
	svg := `<polyline points="9 10 4 15 9 20" /><path d="M20 4v7a4 4 0 0 1-4 4H4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
