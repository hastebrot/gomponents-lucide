package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BadgePoundSterling(children ...g.Node) g.Node {
	svg := `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z" /><path d="M8 12h4" /><path d="M10 16V9.5a2.5 2.5 0 0 1 5 0" /><path d="M8 16h7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
