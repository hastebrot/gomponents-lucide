package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Zap(children ...g.Node) g.Node {
	svg := `<polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
