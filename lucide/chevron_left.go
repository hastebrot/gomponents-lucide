package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronLeft(children ...g.Node) g.Node {
	svg := `<path d="m15 18-6-6 6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
