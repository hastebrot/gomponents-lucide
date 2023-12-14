package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PenTool(children ...g.Node) g.Node {
	svg := `<path d="m12 19 7-7 3 3-7 7-3-3z" /><path d="m18 13-1.5-7.5L2 2l3.5 14.5L13 18l5-5z" /><path d="m2 2 7.586 7.586" /><circle cx="11" cy="11" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
