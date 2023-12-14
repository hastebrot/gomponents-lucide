package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BookmarkCheck(children ...g.Node) g.Node {
	svg := `<path d="m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z" /><path d="m9 10 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
