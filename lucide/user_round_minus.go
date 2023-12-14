package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UserRoundMinus(children ...g.Node) g.Node {
	svg := `<path d="M2 21a8 8 0 0 1 13.292-6" /><circle cx="10" cy="8" r="5" /><path d="M22 19h-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
