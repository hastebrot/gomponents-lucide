package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Reply(children ...g.Node) g.Node {
	svg := `<polyline points="9 17 4 12 9 7" /><path d="M20 18v-2a4 4 0 0 0-4-4H4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
