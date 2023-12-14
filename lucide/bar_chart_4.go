package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BarChart4(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><path d="M13 17V9" /><path d="M18 17V5" /><path d="M8 17v-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
