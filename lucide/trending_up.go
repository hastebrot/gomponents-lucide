package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TrendingUp(children ...g.Node) g.Node {
	svg := `<polyline points="22 7 13.5 15.5 8.5 10.5 2 17" /><polyline points="16 7 22 7 22 13" />`
	return Icon(g.Group(children), g.Raw(svg))
}
