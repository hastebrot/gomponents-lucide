package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FlagTriangleLeft(children ...g.Node) g.Node {
	svg := `<path d="M17 22V2L7 7l10 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
