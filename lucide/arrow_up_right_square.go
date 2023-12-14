package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpRightSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M8 8h8v8" /><path d="m8 16 8-8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
