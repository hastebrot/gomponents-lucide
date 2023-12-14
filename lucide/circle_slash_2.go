package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CircleSlash2(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M22 2 2 22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
