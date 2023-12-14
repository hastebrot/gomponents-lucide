package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MousePointerSquare(children ...g.Node) g.Node {
	svg := `<path d="M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6" /><path d="m12 12 4 10 1.7-4.3L22 16Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
