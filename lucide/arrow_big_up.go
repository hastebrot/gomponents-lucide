package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigUp(children ...g.Node) g.Node {
	svg := `<path d="M9 18v-6H5l7-7 7 7h-4v6H9z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
