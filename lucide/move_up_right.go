package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveUpRight(children ...g.Node) g.Node {
	svg := `<path d="M13 5H19V11" /><path d="M19 5L5 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
