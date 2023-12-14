package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownFromLine(children ...g.Node) g.Node {
	svg := `<path d="M19 3H5" /><path d="M12 21V7" /><path d="m6 15 6 6 6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
