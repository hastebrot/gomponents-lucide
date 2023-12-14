package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Gem(children ...g.Node) g.Node {
	svg := `<path d="M6 3h12l4 6-10 13L2 9Z" /><path d="M11 3 8 9l4 13 4-13-3-6" /><path d="M2 9h20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
