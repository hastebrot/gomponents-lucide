package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveHorizontal(children ...g.Node) g.Node {
	svg := `<polyline points="18 8 22 12 18 16" /><polyline points="6 8 2 12 6 16" /><line x1="2" x2="22" y1="12" y2="12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
