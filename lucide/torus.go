package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Torus(children ...g.Node) g.Node {
	svg := `<ellipse cx="12" cy="11" rx="3" ry="2" /><ellipse cx="12" cy="12.5" rx="10" ry="8.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
