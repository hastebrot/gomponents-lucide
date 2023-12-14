package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Pause(children ...g.Node) g.Node {
	svg := `<rect width="4" height="16" x="6" y="4" /><rect width="4" height="16" x="14" y="4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
