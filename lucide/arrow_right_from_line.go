package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowRightFromLine(children ...g.Node) g.Node {
	svg := `<path d="M3 5v14" /><path d="M21 12H7" /><path d="m15 18 6-6-6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
