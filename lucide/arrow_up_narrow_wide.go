package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpNarrowWide(children ...g.Node) g.Node {
	svg := `<path d="m3 8 4-4 4 4" /><path d="M7 4v16" /><path d="M11 12h4" /><path d="M11 16h7" /><path d="M11 20h10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
