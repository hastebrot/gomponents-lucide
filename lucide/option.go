package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Option(children ...g.Node) g.Node {
	svg := `<path d="M3 3h6l6 18h6" /><path d="M14 3h7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
