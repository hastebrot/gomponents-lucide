package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileCode2(children ...g.Node) g.Node {
	svg := `<path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4" /><polyline points="14 2 14 8 20 8" /><path d="m9 18 3-3-3-3" /><path d="m5 12-3 3 3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
