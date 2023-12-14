package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveDiagonal(children ...g.Node) g.Node {
	svg := `<polyline points="13 5 19 5 19 11" /><polyline points="11 19 5 19 5 13" /><line x1="19" x2="5" y1="5" y2="19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
