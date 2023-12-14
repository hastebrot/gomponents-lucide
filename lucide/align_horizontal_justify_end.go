package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignHorizontalJustifyEnd(children ...g.Node) g.Node {
	svg := `<rect width="6" height="14" x="2" y="5" rx="2" /><rect width="6" height="10" x="12" y="7" rx="2" /><path d="M22 2v20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
