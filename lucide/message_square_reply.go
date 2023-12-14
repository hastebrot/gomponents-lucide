package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareReply(children ...g.Node) g.Node {
	svg := `<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z" /><path d="m10 7-3 3 3 3" /><path d="M17 13v-1a2 2 0 0 0-2-2H7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
