package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func X(children ...g.Node) g.Node {
	svg := `<path d="M18 6 6 18" /><path d="m6 6 12 12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
