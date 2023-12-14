package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlarmClockCheck(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="13" r="8" /><path d="M5 3 2 6" /><path d="m22 6-3-3" /><path d="M6.38 18.7 4 21" /><path d="M17.64 18.67 20 21" /><path d="m9 13 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
