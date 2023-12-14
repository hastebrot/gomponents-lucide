package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StepForward(children ...g.Node) g.Node {
	svg := `<line x1="6" x2="6" y1="4" y2="20" /><polygon points="10,4 20,12 10,20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
