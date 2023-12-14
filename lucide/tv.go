package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tv(children ...g.Node) g.Node {
	svg := `<rect width="20" height="15" x="2" y="7" rx="2" ry="2" /><polyline points="17 2 12 7 7 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
