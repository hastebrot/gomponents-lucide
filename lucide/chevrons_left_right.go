package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsLeftRight(children ...g.Node) g.Node {
	svg := `<path d="m9 7-5 5 5 5" /><path d="m15 7 5 5-5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
