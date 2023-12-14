package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SquareDashedBottomCode(children ...g.Node) g.Node {
	svg := `<path d="m10 10-2 2 2 2" /><path d="m14 14 2-2-2-2" /><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2" /><path d="M9 21h1" /><path d="M14 21h1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
