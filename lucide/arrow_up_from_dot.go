package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpFromDot(children ...g.Node) g.Node {
	svg := `<path d="m5 9 7-7 7 7" /><path d="M12 16V2" /><circle cx="12" cy="21" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
