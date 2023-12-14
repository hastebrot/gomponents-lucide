package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BookmarkX(children ...g.Node) g.Node {
	svg := `<path d="m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z" /><path d="m14.5 7.5-5 5" /><path d="m9.5 7.5 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
