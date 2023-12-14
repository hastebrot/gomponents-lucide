package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpLeftSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M8 16V8h8" /><path d="M16 16 8 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
