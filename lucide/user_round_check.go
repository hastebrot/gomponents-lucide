package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UserRoundCheck(children ...g.Node) g.Node {
	svg := `<path d="M2 21a8 8 0 0 1 13.292-6" /><circle cx="10" cy="8" r="5" /><path d="m16 19 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
