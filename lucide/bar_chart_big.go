package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BarChartBig(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><rect width="4" height="7" x="7" y="10" rx="1" /><rect width="4" height="12" x="15" y="5" rx="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
