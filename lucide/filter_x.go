package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FilterX(children ...g.Node) g.Node {
	svg := `<path d="M13.013 3H2l8 9.46V19l4 2v-8.54l.9-1.055" /><path d="m22 3-5 5" /><path d="m17 3 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
