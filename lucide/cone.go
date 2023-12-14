package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cone(children ...g.Node) g.Node {
	svg := `<path d="m20.9 18.55-8-15.98a1 1 0 0 0-1.8 0l-8 15.98" /><ellipse cx="12" cy="19" rx="9" ry="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
