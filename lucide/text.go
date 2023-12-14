package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Text(children ...g.Node) g.Node {
	svg := `<path d="M17 6.1H3" /><path d="M21 12.1H3" /><path d="M15.1 18H3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
