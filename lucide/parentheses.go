package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Parentheses(children ...g.Node) g.Node {
	svg := `<path d="M8 21s-4-3-4-9 4-9 4-9" /><path d="M16 3s4 3 4 9-4 9-4 9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
