package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func EqualNot(children ...g.Node) g.Node {
	svg := `<line x1="5" x2="19" y1="9" y2="9" /><line x1="5" x2="19" y1="15" y2="15" /><line x1="19" x2="5" y1="5" y2="19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
