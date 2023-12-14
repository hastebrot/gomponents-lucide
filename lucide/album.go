package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Album(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><polyline points="11 3 11 11 14 8 17 11 17 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
