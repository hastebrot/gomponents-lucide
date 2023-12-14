package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GanttChart(children ...g.Node) g.Node {
	svg := `<path d="M8 6h10" /><path d="M6 12h9" /><path d="M11 18h7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
