package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FlipVertical2(children ...g.Node) g.Node {
	svg := `<path d="m17 3-5 5-5-5h10" /><path d="m17 21-5-5-5 5h10" /><path d="M4 12H2" /><path d="M10 12H8" /><path d="M16 12h-2" /><path d="M22 12h-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
