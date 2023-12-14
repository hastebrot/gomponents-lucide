package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareX(children ...g.Node) g.Node {
	svg := `<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z" /><path d="m14.5 7.5-5 5" /><path d="m9.5 7.5 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
