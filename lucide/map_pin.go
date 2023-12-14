package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MapPin(children ...g.Node) g.Node {
	svg := `<path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z" /><circle cx="12" cy="10" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
