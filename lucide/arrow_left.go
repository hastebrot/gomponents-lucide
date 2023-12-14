package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowLeft(children ...g.Node) g.Node {
	svg := `<path d="m12 19-7-7 7-7" /><path d="M19 12H5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
