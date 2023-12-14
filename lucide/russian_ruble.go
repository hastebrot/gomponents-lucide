package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RussianRuble(children ...g.Node) g.Node {
	svg := `<path d="M6 11h8a4 4 0 0 0 0-8H9v18" /><path d="M6 15h8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
