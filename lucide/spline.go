package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Spline(children ...g.Node) g.Node {
	svg := `<circle cx="19" cy="5" r="2" /><circle cx="5" cy="19" r="2" /><path d="M5 17A12 12 0 0 1 17 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
