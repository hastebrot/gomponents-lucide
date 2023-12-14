package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Sailboat(children ...g.Node) g.Node {
	svg := `<path d="M22 18H2a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4Z" /><path d="M21 14 10 2 3 14h18Z" /><path d="M10 2v16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
