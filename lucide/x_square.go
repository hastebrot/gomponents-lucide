package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func XSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><path d="m15 9-6 6" /><path d="m9 9 6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
