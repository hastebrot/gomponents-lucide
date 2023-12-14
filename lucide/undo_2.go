package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Undo2(children ...g.Node) g.Node {
	svg := `<path d="M9 14 4 9l5-5" /><path d="M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5v0a5.5 5.5 0 0 1-5.5 5.5H11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
