package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ungroup(children ...g.Node) g.Node {
	svg := `<rect width="8" height="6" x="5" y="4" rx="1" /><rect width="8" height="6" x="11" y="14" rx="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
