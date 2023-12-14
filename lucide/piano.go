package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Piano(children ...g.Node) g.Node {
	svg := `<path d="M18.5 8c-1.4 0-2.6-.8-3.2-2A6.87 6.87 0 0 0 2 9v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8.5C22 9.6 20.4 8 18.5 8" /><path d="M2 14h20" /><path d="M6 14v4" /><path d="M10 14v4" /><path d="M14 14v4" /><path d="M18 14v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
