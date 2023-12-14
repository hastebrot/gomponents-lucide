package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleOff(children ...g.Node) g.Node {
	svg := `<path d="m2 2 20 20" /><path d="M8.35 2.69A10 10 0 0 1 21.3 15.65" /><path d="M19.08 19.08A10 10 0 1 1 4.92 4.92" />`
	return Icon(g.Group(children), g.Raw(svg))
}
