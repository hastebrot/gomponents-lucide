package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUp10(children ...g.Node) g.Node {
	svg := `<path d="m3 8 4-4 4 4" /><path d="M7 4v16" /><path d="M17 10V4h-2" /><path d="M15 10h4" /><rect x="15" y="14" width="4" height="6" ry="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
