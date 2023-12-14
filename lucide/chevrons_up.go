package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronsUp(children ...g.Node) g.Node {
	svg := `<path d="m17 11-5-5-5 5" /><path d="m17 18-5-5-5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
