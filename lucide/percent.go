package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Percent(children ...g.Node) g.Node {
	svg := `<line x1="19" x2="5" y1="5" y2="19" /><circle cx="6.5" cy="6.5" r="2.5" /><circle cx="17.5" cy="17.5" r="2.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
