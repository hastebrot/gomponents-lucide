package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Key(children ...g.Node) g.Node {
	svg := `<circle cx="7.5" cy="15.5" r="5.5" /><path d="m21 2-9.6 9.6" /><path d="m15.5 7.5 3 3L22 7l-3-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
