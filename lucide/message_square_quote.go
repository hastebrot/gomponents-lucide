package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareQuote(children ...g.Node) g.Node {
	svg := `<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z" /><path d="M8 12a2 2 0 0 0 2-2V8H8" /><path d="M14 12a2 2 0 0 0 2-2V8h-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
