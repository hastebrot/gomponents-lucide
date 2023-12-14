package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleUserRound(children ...g.Node) g.Node {
	svg := `<path d="M18 20a6 6 0 0 0-12 0" /><circle cx="12" cy="10" r="4" /><circle cx="12" cy="12" r="10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
