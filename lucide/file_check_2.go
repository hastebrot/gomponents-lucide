package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileCheck2(children ...g.Node) g.Node {
	svg := `<path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4" /><polyline points="14 2 14 8 20 8" /><path d="m3 15 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
