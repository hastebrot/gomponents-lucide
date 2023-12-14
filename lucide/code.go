package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Code(children ...g.Node) g.Node {
	svg := `<polyline points="16 18 22 12 16 6" /><polyline points="8 6 2 12 8 18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
