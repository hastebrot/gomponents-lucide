package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cast(children ...g.Node) g.Node {
	svg := `<path d="M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6" /><path d="M2 12a9 9 0 0 1 8 8" /><path d="M2 16a5 5 0 0 1 4 4" /><line x1="2" x2="2.01" y1="20" y2="20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
