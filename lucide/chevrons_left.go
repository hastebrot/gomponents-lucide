package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsLeft(children ...g.Node) g.Node {
	svg := `<path d="m11 17-5-5 5-5" /><path d="m18 17-5-5 5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
