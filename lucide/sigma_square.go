package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SigmaSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M16 8.9V7H8l4 5-4 5h8v-1.9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
