package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AppWindow(children ...g.Node) g.Node {
	svg := `<rect x="2" y="4" width="20" height="16" rx="2" /><path d="M10 4v4" /><path d="M2 8h20" /><path d="M6 4v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
