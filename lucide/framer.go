package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Framer(children ...g.Node) g.Node {
	svg := `<path d="M5 16V9h14V2H5l14 14h-7m-7 0 7 7v-7m-7 0h7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
