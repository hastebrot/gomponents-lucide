package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpZA(children ...g.Node) g.Node {
	svg := `<path d="m3 8 4-4 4 4" /><path d="M7 4v16" /><path d="M15 4h5l-5 6h5" /><path d="M15 20v-3.5a2.5 2.5 0 0 1 5 0V20" /><path d="M20 18h-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
