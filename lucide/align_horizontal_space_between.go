package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignHorizontalSpaceBetween(children ...g.Node) g.Node {
	svg := `<rect width="6" height="14" x="3" y="5" rx="2" /><rect width="6" height="10" x="15" y="7" rx="2" /><path d="M3 2v20" /><path d="M21 2v20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
