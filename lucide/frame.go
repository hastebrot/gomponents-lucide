package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Frame(children ...g.Node) g.Node {
	svg := `<line x1="22" x2="2" y1="6" y2="6" /><line x1="22" x2="2" y1="18" y2="18" /><line x1="6" x2="6" y1="2" y2="22" /><line x1="18" x2="18" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
