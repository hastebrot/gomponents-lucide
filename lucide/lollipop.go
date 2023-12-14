package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Lollipop(children ...g.Node) g.Node {
	svg := `<circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" /><path d="M11 11a2 2 0 0 0 4 0 4 4 0 0 0-8 0 6 6 0 0 0 12 0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
