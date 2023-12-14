package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronLeftCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m14 16-4-4 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
