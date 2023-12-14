package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tablet(children ...g.Node) g.Node {
	svg := `<rect width="16" height="20" x="4" y="2" rx="2" ry="2" /><line x1="12" x2="12.01" y1="18" y2="18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
