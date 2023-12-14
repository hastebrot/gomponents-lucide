package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Type(children ...g.Node) g.Node {
	svg := `<polyline points="4 7 4 4 20 4 20 7" /><line x1="9" x2="15" y1="20" y2="20" /><line x1="12" x2="12" y1="4" y2="20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
