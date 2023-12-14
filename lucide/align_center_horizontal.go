package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignCenterHorizontal(children ...g.Node) g.Node {
	svg := `<path d="M2 12h20" /><path d="M10 16v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4" /><path d="M10 8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4" /><path d="M20 16v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1" /><path d="M14 8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
