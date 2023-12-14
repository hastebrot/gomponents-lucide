package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Grip(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="5" r="1" /><circle cx="19" cy="5" r="1" /><circle cx="5" cy="5" r="1" /><circle cx="12" cy="12" r="1" /><circle cx="19" cy="12" r="1" /><circle cx="5" cy="12" r="1" /><circle cx="12" cy="19" r="1" /><circle cx="19" cy="19" r="1" /><circle cx="5" cy="19" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
