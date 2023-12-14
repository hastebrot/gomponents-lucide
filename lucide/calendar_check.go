package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CalendarCheck(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="4" rx="2" ry="2" /><line x1="16" x2="16" y1="2" y2="6" /><line x1="8" x2="8" y1="2" y2="6" /><line x1="3" x2="21" y1="10" y2="10" /><path d="m9 16 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
