package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignVerticalJustifyStart(children ...g.Node) g.Node {
	svg := `<rect width="14" height="6" x="5" y="16" rx="2" /><rect width="10" height="6" x="7" y="6" rx="2" /><path d="M2 2h20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
