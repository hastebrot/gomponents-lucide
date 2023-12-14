package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Bed(children ...g.Node) g.Node {
	svg := `<path d="M2 4v16" /><path d="M2 8h18a2 2 0 0 1 2 2v10" /><path d="M2 17h20" /><path d="M6 8v9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
