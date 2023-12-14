package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PanelBottomInactive(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M14 15h1" /><path d="M19 15h2" /><path d="M3 15h2" /><path d="M9 15h1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
