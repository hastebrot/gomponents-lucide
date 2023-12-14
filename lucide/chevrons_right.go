package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsRight(children ...g.Node) g.Node {
	svg := `<path d="m6 17 5-5-5-5" /><path d="m13 17 5-5-5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
