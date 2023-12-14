package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UserRoundX(children ...g.Node) g.Node {
	svg := `<path d="M2 21a8 8 0 0 1 11.873-7" /><circle cx="10" cy="8" r="5" /><path d="m17 17 5 5" /><path d="m22 17-5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
