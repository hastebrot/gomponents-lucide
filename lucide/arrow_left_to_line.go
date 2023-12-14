package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeftToLine(children ...g.Node) g.Node {
	svg := `<path d="M3 19V5" /><path d="m13 6-6 6 6 6" /><path d="M7 12h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
