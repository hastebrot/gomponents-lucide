package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ZoomOut(children ...g.Node) g.Node {
	svg := `<circle cx="11" cy="11" r="8" /><line x1="21" x2="16.65" y1="21" y2="16.65" /><line x1="8" x2="14" y1="11" y2="11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
