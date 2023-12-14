package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpToLine(children ...g.Node) g.Node {
	svg := `<path d="M5 3h14" /><path d="m18 13-6-6-6 6" /><path d="M12 7v14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
