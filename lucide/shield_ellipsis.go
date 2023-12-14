package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ShieldEllipsis(children ...g.Node) g.Node {
	svg := `<path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10" /><path d="M8 11h.01" /><path d="M12 11h.01" /><path d="M16 11h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
