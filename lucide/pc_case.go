package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PcCase(children ...g.Node) g.Node {
	svg := `<rect width="14" height="20" x="5" y="2" rx="2" /><path d="M15 14h.01" /><path d="M9 6h6" /><path d="M9 10h6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
