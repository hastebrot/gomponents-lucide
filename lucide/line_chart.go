package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func LineChart(children ...g.Node) g.Node {
	svg := `<path d="M3 3v18h18" /><path d="m19 9-5 5-4-4-3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
