package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Clock6(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><polyline points="12 6 12 12 12 16.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
