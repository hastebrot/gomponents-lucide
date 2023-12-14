package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpRightFromCircle(children ...g.Node) g.Node {
	svg := `<path d="M22 12A10 10 0 1 1 12 2" /><path d="M22 2 12 12" /><path d="M16 2h6v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
