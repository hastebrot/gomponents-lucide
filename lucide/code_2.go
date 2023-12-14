package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Code2(children ...g.Node) g.Node {
	svg := `<path d="m18 16 4-4-4-4" /><path d="m6 8-4 4 4 4" /><path d="m14.5 4-5 16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
