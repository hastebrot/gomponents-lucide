package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FlagTriangleRight(children ...g.Node) g.Node {
	svg := `<path d="M7 22V2l10 5-10 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
