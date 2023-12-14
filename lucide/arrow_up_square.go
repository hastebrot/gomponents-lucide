package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="m16 12-4-4-4 4" /><path d="M12 16V8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
