package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Search(children ...g.Node) g.Node {
	svg := `<circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
