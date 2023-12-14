package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveRight(children ...g.Node) g.Node {
	svg := `<path d="M18 8L22 12L18 16" /><path d="M2 12H22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
