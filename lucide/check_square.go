package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CheckSquare(children ...g.Node) g.Node {
	svg := `<path d="m9 11 3 3L22 4" /><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
