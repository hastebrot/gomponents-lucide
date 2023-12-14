package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageCircleWarning(children ...g.Node) g.Node {
	svg := `<path d="M7.9 20A9 9 0 1 0 4 16.1L2 22Z" /><path d="M12 8v4" /><path d="M12 16h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
