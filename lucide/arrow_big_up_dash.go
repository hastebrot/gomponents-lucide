package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigUpDash(children ...g.Node) g.Node {
	svg := `<path d="M9 19h6" /><path d="M9 15v-3H5l7-7 7 7h-4v3H9z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
