package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Maximize2(children ...g.Node) g.Node {
	svg := `<polyline points="15 3 21 3 21 9" /><polyline points="9 21 3 21 3 15" /><line x1="21" x2="14" y1="3" y2="10" /><line x1="3" x2="10" y1="21" y2="14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
