package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignStartHorizontal(children ...g.Node) g.Node {
	svg := `<rect width="6" height="16" x="4" y="6" rx="2" /><rect width="6" height="9" x="14" y="6" rx="2" /><path d="M22 2H2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
