package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeftFromLine(children ...g.Node) g.Node {
	svg := `<path d="m9 6-6 6 6 6" /><path d="M3 12h14" /><path d="M21 19V5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
