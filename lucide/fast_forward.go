package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FastForward(children ...g.Node) g.Node {
	svg := `<polygon points="13 19 22 12 13 5 13 19" /><polygon points="2 19 11 12 2 5 2 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
