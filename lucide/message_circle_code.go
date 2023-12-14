package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageCircleCode(children ...g.Node) g.Node {
	svg := `<path d="M7.9 20A9 9 0 1 0 4 16.1L2 22Z" /><path d="m10 10-2 2 2 2" /><path d="m14 10 2 2-2 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
