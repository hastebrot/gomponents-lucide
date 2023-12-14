package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveDownRight(children ...g.Node) g.Node {
	svg := `<path d="M19 13V19H13" /><path d="M5 5L19 19" />`
	return Icon(g.Group(children), g.Raw(svg))
}
