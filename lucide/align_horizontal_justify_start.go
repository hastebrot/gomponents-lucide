package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignHorizontalJustifyStart(children ...g.Node) g.Node {
	svg := `<rect width="6" height="14" x="6" y="5" rx="2" /><rect width="6" height="10" x="16" y="7" rx="2" /><path d="M2 2v20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
