package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func IterationCcw(children ...g.Node) g.Node {
	svg := `<path d="M20 10c0-4.4-3.6-8-8-8s-8 3.6-8 8 3.6 8 8 8h8" /><polyline points="16 14 20 18 16 22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
