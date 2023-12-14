package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileAxis3d(children ...g.Node) g.Node {
	svg := `<path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" /><polyline points="14 2 14 8 20 8" /><path d="M8 10v8h8" /><path d="m8 18 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}