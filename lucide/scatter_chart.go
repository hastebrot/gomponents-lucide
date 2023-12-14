package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ScatterChart(children ...g.Node) g.Node {
	svg := `<circle cx="7.5" cy="7.5" r=".5" /><circle cx="18.5" cy="5.5" r=".5" /><circle cx="11.5" cy="11.5" r=".5" /><circle cx="7.5" cy="16.5" r=".5" /><circle cx="17.5" cy="14.5" r=".5" /><path d="M3 3v18h18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
