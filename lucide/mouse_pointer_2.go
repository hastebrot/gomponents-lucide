package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MousePointer2(children ...g.Node) g.Node {
	svg := `<path d="m4 4 7.07 17 2.51-7.39L21 11.07z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
