package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleEqual(children ...g.Node) g.Node {
	svg := `<path d="M7 10h10" /><path d="M7 14h10" /><circle cx="12" cy="12" r="10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
