package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUp01(children ...g.Node) g.Node {
	svg := `<path d="m3 8 4-4 4 4" /><path d="M7 4v16" /><rect x="15" y="4" width="4" height="6" ry="2" /><path d="M17 20v-6h-2" /><path d="M15 20h4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
