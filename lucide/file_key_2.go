package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileKey2(children ...g.Node) g.Node {
	svg := `<path d="M4 10V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4" /><polyline points="14 2 14 8 20 8" /><circle cx="4" cy="16" r="2" /><path d="m10 10-4.5 4.5" /><path d="m9 11 1 1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
