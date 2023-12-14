package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Send(children ...g.Node) g.Node {
	svg := `<path d="m22 2-7 20-4-9-9-4Z" /><path d="M22 2 11 13" />`
	return Icon(g.Group(children), g.Raw(svg))
}
