package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Donut(children ...g.Node) g.Node {
	svg := `<path d="M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4 10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3" /><circle cx="12" cy="12" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
