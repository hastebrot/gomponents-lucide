package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignLeft(children ...g.Node) g.Node {
	svg := `<line x1="21" x2="3" y1="6" y2="6" /><line x1="15" x2="3" y1="12" y2="12" /><line x1="17" x2="3" y1="18" y2="18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
