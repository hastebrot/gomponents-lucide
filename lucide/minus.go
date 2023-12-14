package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Minus(children ...g.Node) g.Node {
	svg := `<path d="M5 12h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
