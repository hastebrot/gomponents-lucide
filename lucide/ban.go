package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ban(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m4.9 4.9 14.2 14.2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
