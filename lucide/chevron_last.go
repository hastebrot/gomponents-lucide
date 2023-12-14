package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronLast(children ...g.Node) g.Node {
	svg := `<path d="m7 18 6-6-6-6" /><path d="M17 6v12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
