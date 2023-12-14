package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StretchVertical(children ...g.Node) g.Node {
	svg := `<rect width="6" height="20" x="4" y="2" rx="2" /><rect width="6" height="20" x="14" y="2" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
