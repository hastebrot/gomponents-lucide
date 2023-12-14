package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Map(children ...g.Node) g.Node {
	svg := `<polygon points="3 6 9 3 15 6 21 3 21 18 15 21 9 18 3 21" /><line x1="9" x2="9" y1="3" y2="18" /><line x1="15" x2="15" y1="6" y2="21" />`
	return Icon(g.Group(children), g.Raw(svg))
}
