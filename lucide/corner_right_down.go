package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CornerRightDown(children ...g.Node) g.Node {
	svg := `<polyline points="10 15 15 20 20 15" /><path d="M4 4h7a4 4 0 0 1 4 4v12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
