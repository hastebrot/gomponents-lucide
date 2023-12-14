package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SendHorizontal(children ...g.Node) g.Node {
	svg := `<path d="m3 3 3 9-3 9 19-9Z" /><path d="M6 12h16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
