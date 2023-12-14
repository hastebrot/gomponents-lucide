package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tangent(children ...g.Node) g.Node {
	svg := `<circle cx="17" cy="4" r="2" /><path d="M15.59 5.41 5.41 15.59" /><circle cx="4" cy="17" r="2" /><path d="M12 22s-4-9-1.5-11.5S22 12 22 12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
