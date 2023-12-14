package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Thermometer(children ...g.Node) g.Node {
	svg := `<path d="M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
