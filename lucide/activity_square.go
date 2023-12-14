package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ActivitySquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M17 12h-2l-2 5-2-10-2 5H7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
