package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsDown(children ...g.Node) g.Node {
	svg := `<path d="m7 6 5 5 5-5" /><path d="m7 13 5 5 5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
