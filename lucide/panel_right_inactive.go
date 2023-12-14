package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PanelRightInactive(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M15 14v1" /><path d="M15 19v2" /><path d="M15 3v2" /><path d="M15 9v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
