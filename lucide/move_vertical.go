package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveVertical(children ...g.Node) g.Node {
	svg := `<polyline points="8 18 12 22 16 18" /><polyline points="8 6 12 2 16 6" /><line x1="12" x2="12" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
