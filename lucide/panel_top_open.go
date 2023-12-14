package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PanelTopOpen(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><line x1="3" x2="21" y1="9" y2="9" /><path d="m15 14-3 3-3-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
