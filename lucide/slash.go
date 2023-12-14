package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Slash(children ...g.Node) g.Node {
	svg := `<path d="M22 2 2 22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
