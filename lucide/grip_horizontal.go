package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GripHorizontal(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="9" r="1" /><circle cx="19" cy="9" r="1" /><circle cx="5" cy="9" r="1" /><circle cx="12" cy="15" r="1" /><circle cx="19" cy="15" r="1" /><circle cx="5" cy="15" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
