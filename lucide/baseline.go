package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Baseline(children ...g.Node) g.Node {
	svg := `<path d="M4 20h16" /><path d="m6 16 6-12 6 12" /><path d="M8 12h8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
