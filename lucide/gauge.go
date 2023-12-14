package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Gauge(children ...g.Node) g.Node {
	svg := `<path d="m12 14 4-4" /><path d="M3.34 19a10 10 0 1 1 17.32 0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
