package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileX2(children ...g.Node) g.Node {
	svg := `<path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4" /><path d="M14 2v6h6" /><path d="m3 12.5 5 5" /><path d="m8 12.5-5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
