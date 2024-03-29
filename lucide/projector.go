package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Projector(children ...g.Node) g.Node {
	svg := `<path d="M5 7 3 5" /><path d="M9 6V3" /><path d="m13 7 2-2" /><circle cx="9" cy="13" r="3" /><path d="M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17" /><path d="M16 16h2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
