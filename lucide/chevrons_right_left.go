package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsRightLeft(children ...g.Node) g.Node {
	svg := `<path d="m20 17-5-5 5-5" /><path d="m4 17 5-5-5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
