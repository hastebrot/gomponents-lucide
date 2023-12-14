package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Club(children ...g.Node) g.Node {
	svg := `<path d="M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6Z" /><path d="M12 17.66L12 22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
