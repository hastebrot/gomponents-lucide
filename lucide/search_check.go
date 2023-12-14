package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SearchCheck(children ...g.Node) g.Node {
	svg := `<path d="m8 11 2 2 4-4" /><circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
