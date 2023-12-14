package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerLeftUp(children ...g.Node) g.Node {
	svg := `<polyline points="14 9 9 4 4 9" /><path d="M20 20h-7a4 4 0 0 1-4-4V4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
