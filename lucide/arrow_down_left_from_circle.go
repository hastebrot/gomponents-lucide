package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownLeftFromCircle(children ...g.Node) g.Node {
	svg := `<path d="M2 12a10 10 0 1 1 10 10" /><path d="m2 22 10-10" /><path d="M8 22H2v-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
