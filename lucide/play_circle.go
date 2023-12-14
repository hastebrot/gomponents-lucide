package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PlayCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><polygon points="10 8 16 12 10 16 10 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
