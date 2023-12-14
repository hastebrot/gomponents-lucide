package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Check(children ...g.Node) g.Node {
	svg := `<path d="M20 6 9 17l-5-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
