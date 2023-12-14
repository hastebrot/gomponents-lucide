package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronUp(children ...g.Node) g.Node {
	svg := `<path d="m18 15-6-6-6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
