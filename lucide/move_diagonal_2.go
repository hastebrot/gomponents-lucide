package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveDiagonal2(children ...g.Node) g.Node {
	svg := `<polyline points="5 11 5 5 11 5" /><polyline points="19 13 19 19 13 19" /><line x1="5" x2="19" y1="5" y2="19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
