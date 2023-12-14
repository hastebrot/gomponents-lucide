package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PersonStanding(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="5" r="1" /><path d="m9 20 3-6 3 6" /><path d="m6 8 6 2 6-2" /><path d="M12 10v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
