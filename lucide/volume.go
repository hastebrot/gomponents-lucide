package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Volume(children ...g.Node) g.Node {
	svg := `<polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
