package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Pilcrow(children ...g.Node) g.Node {
	svg := `<path d="M13 4v16" /><path d="M17 4v16" /><path d="M19 4H9.5a4.5 4.5 0 0 0 0 9H13" />`
	return Icon(g.Group(children), g.Raw(svg))
}
