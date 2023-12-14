package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Sigma(children ...g.Node) g.Node {
	svg := `<path d="M18 7V4H6l6 8-6 8h12v-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
