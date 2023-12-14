package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cylinder(children ...g.Node) g.Node {
	svg := `<ellipse cx="12" cy="5" rx="9" ry="3" /><path d="M3 5v14a9 3 0 0 0 18 0V5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
