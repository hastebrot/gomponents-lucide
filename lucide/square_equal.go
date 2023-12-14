package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SquareEqual(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M7 10h10" /><path d="M7 14h10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
