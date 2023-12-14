package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Filter(children ...g.Node) g.Node {
	svg := `<polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
