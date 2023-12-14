package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PowerSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M12 7v5" /><path d="M8 9a5.14 5.14 0 0 0 4 8 4.95 4.95 0 0 0 4-8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
