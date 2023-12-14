package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowRight(children ...g.Node) g.Node {
	svg := `<path d="M5 12h14" /><path d="m12 5 7 7-7 7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
