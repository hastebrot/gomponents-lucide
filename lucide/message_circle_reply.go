package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageCircleReply(children ...g.Node) g.Node {
	svg := `<path d="M7.9 20A9 9 0 1 0 4 16.1L2 22Z" /><path d="m10 15-3-3 3-3" /><path d="M7 12h7a2 2 0 0 1 2 2v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
