package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Link2Off(children ...g.Node) g.Node {
	svg := `<path d="M9 17H7A5 5 0 0 1 7 7" /><path d="M15 7h2a5 5 0 0 1 4 8" /><line x1="8" x2="12" y1="12" y2="12" /><line x1="2" x2="22" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
