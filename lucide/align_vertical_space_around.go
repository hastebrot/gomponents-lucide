package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignVerticalSpaceAround(children ...g.Node) g.Node {
	svg := `<rect width="10" height="6" x="7" y="9" rx="2" /><path d="M22 20H2" /><path d="M22 4H2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
