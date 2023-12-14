package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleSlash(children ...g.Node) g.Node {
	svg := `<line x1="9" x2="15" y1="15" y2="9" /><circle cx="12" cy="12" r="10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
