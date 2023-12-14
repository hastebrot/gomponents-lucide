package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Navigation(children ...g.Node) g.Node {
	svg := `<polygon points="3 11 22 2 13 21 11 13 3 11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
