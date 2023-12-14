package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CreditCard(children ...g.Node) g.Node {
	svg := `<rect width="20" height="14" x="2" y="5" rx="2" /><line x1="2" x2="22" y1="10" y2="10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
