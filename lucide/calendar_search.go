package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CalendarSearch(children ...g.Node) g.Node {
	svg := `<path d="M21 12V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7.5" /><path d="M16 2v4" /><path d="M8 2v4" /><path d="M3 10h18" /><path d="M18 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6v0Z" /><path d="m22 22-1.5-1.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
