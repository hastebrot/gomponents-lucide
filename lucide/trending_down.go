package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TrendingDown(children ...g.Node) g.Node {
	svg := `<polyline points="22 17 13.5 8.5 8.5 13.5 2 7" /><polyline points="16 17 22 17 22 11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
