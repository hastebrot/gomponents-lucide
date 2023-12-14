package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Mountain(children ...g.Node) g.Node {
	svg := `<path d="m8 3 4 8 5-5 5 15H2L8 3z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
