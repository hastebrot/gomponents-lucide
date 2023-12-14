package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListStart(children ...g.Node) g.Node {
	svg := `<path d="M16 12H3" /><path d="M16 18H3" /><path d="M10 6H3" /><path d="M21 18V8a2 2 0 0 0-2-2h-5" /><path d="m16 8-2-2 2-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
