package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Dot(children ...g.Node) g.Node {
	svg := `<circle cx="12.1" cy="12.1" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
