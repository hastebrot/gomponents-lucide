package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SwissFranc(children ...g.Node) g.Node {
	svg := `<path d="M10 21V3h8" /><path d="M6 16h9" /><path d="M10 9.5h7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
