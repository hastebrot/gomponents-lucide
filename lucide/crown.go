package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Crown(children ...g.Node) g.Node {
	svg := `<path d="m2 4 3 12h14l3-12-6 7-4-7-4 7-6-7zm3 16h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
