package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronRight(children ...g.Node) g.Node {
	svg := `<path d="m9 18 6-6-6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
