package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RailSymbol(children ...g.Node) g.Node {
	svg := `<path d="M5 15h14" /><path d="M5 9h14" /><path d="m14 20-5-5 6-6-5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
