package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Refrigerator(children ...g.Node) g.Node {
	svg := `<path d="M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Z" /><path d="M5 10h14" /><path d="M15 7v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
