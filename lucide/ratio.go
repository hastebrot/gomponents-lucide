package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ratio(children ...g.Node) g.Node {
	svg := `<rect width="12" height="20" x="6" y="2" rx="2" /><rect width="20" height="12" x="2" y="6" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
