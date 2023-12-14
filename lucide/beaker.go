package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Beaker(children ...g.Node) g.Node {
	svg := `<path d="M4.5 3h15" /><path d="M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3" /><path d="M6 14h12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
