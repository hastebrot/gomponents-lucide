package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func View(children ...g.Node) g.Node {
	svg := `<path d="M5 12s2.545-5 7-5c4.454 0 7 5 7 5s-2.546 5-7 5c-4.455 0-7-5-7-5z" /><path d="M12 13a1 1 0 1 0 0-2 1 1 0 0 0 0 2z" /><path d="M21 17v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2" /><path d="M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
