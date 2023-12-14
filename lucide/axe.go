package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Axe(children ...g.Node) g.Node {
	svg := `<path d="m14 12-8.5 8.5a2.12 2.12 0 1 1-3-3L11 9" /><path d="M15 13 9 7l4-4 6 6h3a8 8 0 0 1-7 7z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
