package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func VenetianMask(children ...g.Node) g.Node {
	svg := `<path d="M2 12a5 5 0 0 0 5 5 8 8 0 0 1 5 2 8 8 0 0 1 5-2 5 5 0 0 0 5-5V7h-5a8 8 0 0 0-5 2 8 8 0 0 0-5-2H2Z" /><path d="M6 11c1.5 0 3 .5 3 2-2 0-3 0-3-2Z" /><path d="M18 11c-1.5 0-3 .5-3 2 2 0 3 0 3-2Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
