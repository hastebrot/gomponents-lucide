package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Terminal(children ...g.Node) g.Node {
	svg := `<polyline points="4 17 10 11 4 5" /><line x1="12" x2="20" y1="19" y2="19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
