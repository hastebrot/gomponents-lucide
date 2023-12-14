package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Divide(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="6" r="1" /><line x1="5" x2="19" y1="12" y2="12" /><circle cx="12" cy="18" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
