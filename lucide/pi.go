package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Pi(children ...g.Node) g.Node {
	svg := `<line x1="9" x2="9" y1="4" y2="20" /><path d="M4 7c0-1.7 1.3-3 3-3h13" /><path d="M18 20c-1.7 0-3-1.3-3-3V4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
