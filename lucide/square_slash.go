package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SquareSlash(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><line x1="9" x2="15" y1="15" y2="9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
