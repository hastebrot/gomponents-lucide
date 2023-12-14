package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Meh(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><line x1="8" x2="16" y1="15" y2="15" /><line x1="9" x2="9.01" y1="9" y2="9" /><line x1="15" x2="15.01" y1="9" y2="9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
