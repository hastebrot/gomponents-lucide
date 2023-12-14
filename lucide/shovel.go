package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Shovel(children ...g.Node) g.Node {
	svg := `<path d="M2 22v-5l5-5 5 5-5 5z" /><path d="M9.5 14.5 16 8" /><path d="m17 2 5 5-.5.5a3.53 3.53 0 0 1-5 0s0 0 0 0a3.53 3.53 0 0 1 0-5L17 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
