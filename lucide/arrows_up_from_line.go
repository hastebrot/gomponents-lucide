package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowsUpFromLine(children ...g.Node) g.Node {
	svg := `<path d="m4 6 3-3 3 3" /><path d="M7 17V3" /><path d="m14 6 3-3 3 3" /><path d="M17 17V3" /><path d="M4 21h16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
