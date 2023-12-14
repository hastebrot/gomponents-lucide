package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerRightUp(children ...g.Node) g.Node {
	svg := `<polyline points="10 9 15 4 20 9" /><path d="M4 20h7a4 4 0 0 0 4-4V4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
