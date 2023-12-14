package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SpellCheck2(children ...g.Node) g.Node {
	svg := `<path d="m6 16 6-12 6 12" /><path d="M8 12h8" /><path d="M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
