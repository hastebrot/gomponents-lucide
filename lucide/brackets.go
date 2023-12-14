package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Brackets(children ...g.Node) g.Node {
	svg := `<path d="M16 3h3v18h-3" /><path d="M8 21H5V3h3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
