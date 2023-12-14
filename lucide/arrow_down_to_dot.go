package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDownToDot(children ...g.Node) g.Node {
	svg := `<path d="M12 2v14" /><path d="m19 9-7 7-7-7" /><circle cx="12" cy="21" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
