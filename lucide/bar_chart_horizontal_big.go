package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BarChartHorizontalBig(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><rect width="12" height="4" x="7" y="5" rx="1" /><rect width="7" height="4" x="7" y="13" rx="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
