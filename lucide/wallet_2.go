package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Wallet2(children ...g.Node) g.Node {
	svg := `<path d="M17 14h.01" /><path d="M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
