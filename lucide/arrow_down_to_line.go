package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownToLine(children ...g.Node) g.Node {
	svg := `<path d="M12 17V3" /><path d="m6 11 6 6 6-6" /><path d="M19 21H5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
