package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TextQuote(children ...g.Node) g.Node {
	svg := `<path d="M17 6H3" /><path d="M21 12H8" /><path d="M21 18H8" /><path d="M3 12v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
