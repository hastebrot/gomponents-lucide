package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownWideNarrow(children ...g.Node) g.Node {
	svg := `<path d="m3 16 4 4 4-4" /><path d="M7 20V4" /><path d="M11 4h10" /><path d="M11 8h7" /><path d="M11 12h4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
