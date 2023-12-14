package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BarChartHorizontal(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><path d="M7 16h8" /><path d="M7 11h12" /><path d="M7 6h3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
