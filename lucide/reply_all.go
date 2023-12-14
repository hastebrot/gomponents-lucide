package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ReplyAll(children ...g.Node) g.Node {
	svg := `<polyline points="7 17 2 12 7 7" /><polyline points="12 17 7 12 12 7" /><path d="M22 18v-2a4 4 0 0 0-4-4H7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
