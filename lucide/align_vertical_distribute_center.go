package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignVerticalDistributeCenter(children ...g.Node) g.Node {
	svg := `<rect width="14" height="6" x="5" y="14" rx="2" /><rect width="10" height="6" x="7" y="4" rx="2" /><path d="M22 7h-5" /><path d="M7 7H1" /><path d="M22 17h-3" /><path d="M5 17H2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
