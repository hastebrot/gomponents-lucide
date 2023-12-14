package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SquareUser(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><circle cx="12" cy="10" r="3" /><path d="M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
