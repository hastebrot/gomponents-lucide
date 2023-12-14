package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RedoDot(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="17" r="1" /><path d="M21 7v6h-6" /><path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
