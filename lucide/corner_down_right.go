package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerDownRight(children ...g.Node) g.Node {
	svg := `<polyline points="15 10 20 15 15 20" /><path d="M4 4v7a4 4 0 0 0 4 4h12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
