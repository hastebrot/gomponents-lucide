package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BookmarkPlus(children ...g.Node) g.Node {
	svg := `<path d="m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z" /><line x1="12" x2="12" y1="7" y2="13" /><line x1="15" x2="9" y1="10" y2="10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
