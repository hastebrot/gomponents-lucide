package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsDownUp(children ...g.Node) g.Node {
	svg := `<path d="m7 20 5-5 5 5" /><path d="m7 4 5 5 5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
