package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Anchor(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="5" r="3" /><line x1="12" x2="12" y1="22" y2="8" /><path d="M5 12H2a10 10 0 0 0 20 0h-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
