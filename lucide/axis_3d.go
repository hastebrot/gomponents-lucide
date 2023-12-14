package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Axis3d(children ...g.Node) g.Node {
	svg := `<path d="M4 4v16h16" /><path d="m4 20 7-7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
