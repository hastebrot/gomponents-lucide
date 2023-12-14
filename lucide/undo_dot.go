package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UndoDot(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="17" r="1" /><path d="M3 7v6h6" /><path d="M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13" />`
	return Icon(g.Group(children), g.Raw(svg))
}
