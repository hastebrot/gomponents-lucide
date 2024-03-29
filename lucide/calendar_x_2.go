package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CalendarX2(children ...g.Node) g.Node {
	svg := `<path d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8" /><line x1="16" x2="16" y1="2" y2="6" /><line x1="8" x2="8" y1="2" y2="6" /><line x1="3" x2="21" y1="10" y2="10" /><line x1="17" x2="22" y1="17" y2="22" /><line x1="17" x2="22" y1="22" y2="17" />`
	return Icon(g.Group(children), g.Raw(svg))
}
