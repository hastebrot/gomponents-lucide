package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownNarrowWide(children ...g.Node) g.Node {
	svg := `<path d="m3 16 4 4 4-4" /><path d="M7 20V4" /><path d="M11 4h4" /><path d="M11 8h7" /><path d="M11 12h10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
