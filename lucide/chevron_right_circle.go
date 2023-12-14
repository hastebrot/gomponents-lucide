package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ChevronRightCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m10 8 4 4-4 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
