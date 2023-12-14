package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Heading5(children ...g.Node) g.Node {
	svg := `<path d="M4 12h8" /><path d="M4 18V6" /><path d="M12 18V6" /><path d="M17 13v-3h4" /><path d="M17 17.7c.4.2.8.3 1.3.3 1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17" />`
	return Icon(g.Group(children), g.Raw(svg))
}
