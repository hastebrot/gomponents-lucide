package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignEndVertical(children ...g.Node) g.Node {
	svg := `<rect width="16" height="6" x="2" y="4" rx="2" /><rect width="9" height="6" x="9" y="14" rx="2" /><path d="M22 22V2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
