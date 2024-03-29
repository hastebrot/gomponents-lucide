package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SkipForward(children ...g.Node) g.Node {
	svg := `<polygon points="5 4 15 12 5 20 5 4" /><line x1="19" x2="19" y1="5" y2="19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
