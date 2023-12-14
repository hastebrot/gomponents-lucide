package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	svg := `<path d="M17 7 7 17" /><path d="M17 17H7V7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
