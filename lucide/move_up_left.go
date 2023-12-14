package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveUpLeft(children ...g.Node) g.Node {
	svg := `<path d="M5 11V5H11" /><path d="M5 5L19 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
