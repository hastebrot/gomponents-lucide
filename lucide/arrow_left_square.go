package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeftSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="m12 8-4 4 4 4" /><path d="M16 12H8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
