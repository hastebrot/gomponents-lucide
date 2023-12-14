package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownRightFromCircle(children ...g.Node) g.Node {
	svg := `<path d="M12 22a10 10 0 1 1 10-10" /><path d="M22 22 12 12" /><path d="M22 16v6h-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
