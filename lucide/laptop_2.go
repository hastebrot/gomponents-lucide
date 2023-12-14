package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Laptop2(children ...g.Node) g.Node {
	svg := `<rect width="18" height="12" x="3" y="4" rx="2" ry="2" /><line x1="2" x2="22" y1="20" y2="20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
