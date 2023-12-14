package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RectangleHorizontal(children ...g.Node) g.Node {
	svg := `<rect width="20" height="12" x="2" y="6" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
