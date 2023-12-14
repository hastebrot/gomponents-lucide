package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FishSymbol(children ...g.Node) g.Node {
	svg := `<path d="M2 16s9-15 20-4C11 23 2 8 2 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
