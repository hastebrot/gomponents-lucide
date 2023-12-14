package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Trello(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><rect width="3" height="9" x="7" y="7" /><rect width="3" height="5" x="14" y="7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
