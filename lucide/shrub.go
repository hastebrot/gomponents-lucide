package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Shrub(children ...g.Node) g.Node {
	svg := `<path d="M12 22v-7l-2-2" /><path d="M17 8v.8A6 6 0 0 1 13.8 20v0H10v0A6.5 6.5 0 0 1 7 8h0a5 5 0 0 1 10 0Z" /><path d="m14 14-2 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
