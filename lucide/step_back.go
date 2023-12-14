package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StepBack(children ...g.Node) g.Node {
	svg := `<line x1="18" x2="18" y1="20" y2="4" /><polygon points="14,20 4,12 14,4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
