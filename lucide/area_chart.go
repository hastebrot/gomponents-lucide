package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AreaChart(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><path d="M7 12v5h12V8l-5 5-4-4Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
