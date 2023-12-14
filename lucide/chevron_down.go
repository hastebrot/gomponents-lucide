package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronDown(children ...g.Node) g.Node {
	svg := `<path d="m6 9 6 6 6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
