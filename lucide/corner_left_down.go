package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerLeftDown(children ...g.Node) g.Node {
	svg := `<polyline points="14 15 9 20 4 15" /><path d="M20 4h-7a4 4 0 0 0-4 4v12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
