package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ampersand(children ...g.Node) g.Node {
	svg := `<path d="M17.5 12c0 4.4-3.6 8-8 8A4.5 4.5 0 0 1 5 15.5c0-6 8-4 8-8.5a3 3 0 1 0-6 0c0 3 2.5 8.5 12 13" /><path d="M16 12h3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
