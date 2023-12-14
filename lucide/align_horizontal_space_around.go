package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignHorizontalSpaceAround(children ...g.Node) g.Node {
	svg := `<rect width="6" height="10" x="9" y="7" rx="2" /><path d="M4 22V2" /><path d="M20 22V2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
