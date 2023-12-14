package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigRight(children ...g.Node) g.Node {
	svg := `<path d="M6 9h6V5l7 7-7 7v-4H6V9z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
