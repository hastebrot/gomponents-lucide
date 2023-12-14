package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeftCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M16 12H8" /><path d="m12 8-4 4 4 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
