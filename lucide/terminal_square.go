package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TerminalSquare(children ...g.Node) g.Node {
	svg := `<path d="m7 11 2-2-2-2" /><path d="M11 13h4" /><rect width="18" height="18" x="3" y="3" rx="2" ry="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
