package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Navigation2(children ...g.Node) g.Node {
	svg := `<polygon points="12 2 19 21 12 17 5 21 12 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
