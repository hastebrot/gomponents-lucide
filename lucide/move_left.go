package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveLeft(children ...g.Node) g.Node {
	svg := `<path d="M6 8L2 12L6 16" /><path d="M2 12H22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
