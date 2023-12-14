package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveDownLeft(children ...g.Node) g.Node {
	svg := `<path d="M11 19H5V13" /><path d="M19 5L5 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
