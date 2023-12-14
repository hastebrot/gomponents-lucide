package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignHorizontalDistributeCenter(children ...g.Node) g.Node {
	svg := `<rect width="6" height="14" x="4" y="5" rx="2" /><rect width="6" height="10" x="14" y="7" rx="2" /><path d="M17 22v-5" /><path d="M17 7V2" /><path d="M7 22v-3" /><path d="M7 5V2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
