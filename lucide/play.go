package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Play(children ...g.Node) g.Node {
	svg := `<polygon points="5 3 19 12 5 21 5 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
