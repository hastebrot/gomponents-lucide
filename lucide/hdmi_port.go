package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func HdmiPort(children ...g.Node) g.Node {
	svg := `<path d="M22 9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1l2 2h12l2-2h1a1 1 0 0 0 1-1Z" /><path d="M7.5 12h9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
