package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsUpDown(children ...g.Node) g.Node {
	svg := `<path d="m7 15 5 5 5-5" /><path d="m7 9 5-5 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
