package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Rows(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><line x1="3" x2="21" y1="12" y2="12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
