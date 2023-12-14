package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SkipBack(children ...g.Node) g.Node {
	svg := `<polygon points="19 20 9 12 19 4 19 20" /><line x1="5" x2="5" y1="19" y2="5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
