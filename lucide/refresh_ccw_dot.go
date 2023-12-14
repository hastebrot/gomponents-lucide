package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RefreshCcwDot(children ...g.Node) g.Node {
	svg := `<path d="M3 2v6h6" /><path d="M21 12A9 9 0 0 0 6 5.3L3 8" /><path d="M21 22v-6h-6" /><path d="M3 12a9 9 0 0 0 15 6.7l3-2.7" /><circle cx="12" cy="12" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
