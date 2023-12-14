package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageCircleHeart(children ...g.Node) g.Node {
	svg := `<path d="M7.9 20A9 9 0 1 0 4 16.1L2 22Z" /><path d="M15.8 9.2a2.5 2.5 0 0 0-3.5 0l-.3.4-.35-.3a2.42 2.42 0 1 0-3.2 3.6l3.6 3.5 3.6-3.5c1.2-1.2 1.1-2.7.2-3.7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
