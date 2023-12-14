package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Space(children ...g.Node) g.Node {
	svg := `<path d="M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
