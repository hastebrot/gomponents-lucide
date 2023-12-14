package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MousePointerSquareDashed(children ...g.Node) g.Node {
	svg := `<path d="M5 3a2 2 0 0 0-2 2" /><path d="M19 3a2 2 0 0 1 2 2" /><path d="m12 12 4 10 1.7-4.3L22 16Z" /><path d="M5 21a2 2 0 0 1-2-2" /><path d="M9 3h1" /><path d="M9 21h2" /><path d="M14 3h1" /><path d="M3 9v1" /><path d="M21 9v2" /><path d="M3 14v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
