package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SquareUserRound(children ...g.Node) g.Node {
	svg := `<path d="M18 21a6 6 0 0 0-12 0" /><circle cx="12" cy="11" r="4" /><rect width="18" height="18" x="3" y="3" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
