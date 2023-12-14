package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AlignStartVertical(children ...g.Node) g.Node {
	svg := `<rect width="9" height="6" x="6" y="14" rx="2" /><rect width="16" height="6" x="6" y="4" rx="2" /><path d="M2 2v20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
