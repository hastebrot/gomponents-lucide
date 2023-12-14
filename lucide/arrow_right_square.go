package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowRightSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M8 12h8" /><path d="m12 16 4-4-4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
