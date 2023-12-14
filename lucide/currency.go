package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Currency(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="8" /><line x1="3" x2="6" y1="3" y2="6" /><line x1="21" x2="18" y1="3" y2="6" /><line x1="3" x2="6" y1="21" y2="18" /><line x1="21" x2="18" y1="21" y2="18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
