package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TreeDeciduous(children ...g.Node) g.Node {
	svg := `<path d="M8 19a4 4 0 0 1-2.24-7.32A3.5 3.5 0 0 1 9 6.03V6a3 3 0 1 1 6 0v.04a3.5 3.5 0 0 1 3.24 5.65A4 4 0 0 1 16 19Z" /><path d="M12 19v3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
