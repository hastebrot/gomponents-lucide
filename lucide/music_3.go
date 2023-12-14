package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Music3(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="18" r="4" /><path d="M16 18V2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
