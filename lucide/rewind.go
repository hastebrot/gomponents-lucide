package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Rewind(children ...g.Node) g.Node {
	svg := `<polygon points="11 19 2 12 11 5 11 19" /><polygon points="22 19 13 12 22 5 22 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
