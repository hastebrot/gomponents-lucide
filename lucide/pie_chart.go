package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PieChart(children ...g.Node) g.Node {
	svg := `<path d="M21.21 15.89A10 10 0 1 1 8 2.83" /><path d="M22 12A10 10 0 0 0 12 2v10z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
