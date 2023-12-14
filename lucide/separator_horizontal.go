package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SeparatorHorizontal(children ...g.Node) g.Node {
	svg := `<line x1="3" x2="21" y1="12" y2="12" /><polyline points="8 8 12 4 16 8" /><polyline points="16 16 12 20 8 16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
