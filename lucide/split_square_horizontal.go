package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SplitSquareHorizontal(children ...g.Node) g.Node {
	svg := `<path d="M8 19H5c-1 0-2-1-2-2V7c0-1 1-2 2-2h3" /><path d="M16 5h3c1 0 2 1 2 2v10c0 1-1 2-2 2h-3" /><line x1="12" x2="12" y1="4" y2="20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
