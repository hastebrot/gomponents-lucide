package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PanelLeftOpen(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><path d="M9 3v18" /><path d="m14 9 3 3-3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
