package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleEllipsis(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M17 12h.01" /><path d="M12 12h.01" /><path d="M7 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
