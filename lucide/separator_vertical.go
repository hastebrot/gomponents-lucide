package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SeparatorVertical(children ...g.Node) g.Node {
	svg := `<line x1="12" x2="12" y1="3" y2="21" /><polyline points="8 8 4 12 8 16" /><polyline points="16 16 20 12 16 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
