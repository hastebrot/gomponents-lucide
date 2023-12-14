package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UserRound(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="8" r="5" /><path d="M20 21a8 8 0 0 0-16 0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
