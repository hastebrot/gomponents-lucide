package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownRight(children ...g.Node) g.Node {
	svg := `<path d="m7 7 10 10" /><path d="M17 7v10H7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
