package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignVerticalDistributeStart(children ...g.Node) g.Node {
	svg := `<rect width="14" height="6" x="5" y="14" rx="2" /><rect width="10" height="6" x="7" y="4" rx="2" /><path d="M2 14h20" /><path d="M2 4h20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
