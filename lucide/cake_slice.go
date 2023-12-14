package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CakeSlice(children ...g.Node) g.Node {
	svg := `<circle cx="9" cy="7" r="2" /><path d="M7.2 7.9 3 11v9c0 .6.4 1 1 1h16c.6 0 1-.4 1-1v-9c0-2-3-6-7-8l-3.6 2.6" /><path d="M16 13H3" /><path d="M16 17H3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
