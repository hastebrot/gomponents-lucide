package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListOrdered(children ...g.Node) g.Node {
	svg := `<line x1="10" x2="21" y1="6" y2="6" /><line x1="10" x2="21" y1="12" y2="12" /><line x1="10" x2="21" y1="18" y2="18" /><path d="M4 6h1v4" /><path d="M4 10h2" /><path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
