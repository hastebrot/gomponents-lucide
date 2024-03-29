package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GaugeCircle(children ...g.Node) g.Node {
	svg := `<path d="M15.6 2.7a10 10 0 1 0 5.7 5.7" /><circle cx="12" cy="12" r="2" /><path d="M13.4 10.6 19 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
