package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func IterationCw(children ...g.Node) g.Node {
	svg := `<path d="M4 10c0-4.4 3.6-8 8-8s8 3.6 8 8-3.6 8-8 8H4" /><polyline points="8 22 4 18 8 14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
