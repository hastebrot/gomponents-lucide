package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigRightDash(children ...g.Node) g.Node {
	svg := `<path d="M5 9v6" /><path d="M9 9h3V5l7 7-7 7v-4H9V9z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
