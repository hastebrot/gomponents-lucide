package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Forward(children ...g.Node) g.Node {
	svg := `<polyline points="15 17 20 12 15 7" /><path d="M4 18v-2a4 4 0 0 1 4-4h12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
