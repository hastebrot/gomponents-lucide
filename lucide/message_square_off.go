package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareOff(children ...g.Node) g.Node {
	svg := `<path d="M21 15V5a2 2 0 0 0-2-2H9" /><path d="m2 2 20 20" /><path d="M3.6 3.6c-.4.3-.6.8-.6 1.4v16l4-4h10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
