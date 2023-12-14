package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cloud(children ...g.Node) g.Node {
	svg := `<path d="M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
