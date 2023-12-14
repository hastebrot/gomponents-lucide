package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleDot(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><circle cx="12" cy="12" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
