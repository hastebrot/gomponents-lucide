package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Volume1(children ...g.Node) g.Node {
	svg := `<polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5" /><path d="M15.54 8.46a5 5 0 0 1 0 7.07" />`
	return Icon(g.Group(children), g.Raw(svg))
}
