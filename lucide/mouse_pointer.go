package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MousePointer(children ...g.Node) g.Node {
	svg := `<path d="m3 3 7.07 16.97 2.51-7.39 7.39-2.51L3 3z" /><path d="m13 13 6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
