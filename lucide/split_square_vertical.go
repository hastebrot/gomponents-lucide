package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SplitSquareVertical(children ...g.Node) g.Node {
	svg := `<path d="M5 8V5c0-1 1-2 2-2h10c1 0 2 1 2 2v3" /><path d="M19 16v3c0 1-1 2-2 2H7c-1 0-2-1-2-2v-3" /><line x1="4" x2="20" y1="12" y2="12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
