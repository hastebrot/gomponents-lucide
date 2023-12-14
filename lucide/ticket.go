package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ticket(children ...g.Node) g.Node {
	svg := `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z" /><path d="M13 5v2" /><path d="M13 17v2" /><path d="M13 11v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
