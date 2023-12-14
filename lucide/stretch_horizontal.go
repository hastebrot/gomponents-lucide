package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StretchHorizontal(children ...g.Node) g.Node {
	svg := `<rect width="20" height="6" x="2" y="4" rx="2" /><rect width="20" height="6" x="2" y="14" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
