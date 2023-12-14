package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RotateCcw(children ...g.Node) g.Node {
	svg := `<path d="M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8" /><path d="M3 3v5h5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
