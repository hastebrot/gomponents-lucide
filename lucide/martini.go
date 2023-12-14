package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Martini(children ...g.Node) g.Node {
	svg := `<path d="M8 22h8" /><path d="M12 11v11" /><path d="m19 3-7 8-7-8Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
