package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareShare(children ...g.Node) g.Node {
	svg := `<path d="M21 12v3a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h7" /><path d="M16 3h5v5" /><path d="m16 8 5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
