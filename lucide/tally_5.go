package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tally5(children ...g.Node) g.Node {
	svg := `<path d="M4 4v16" /><path d="M9 4v16" /><path d="M14 4v16" /><path d="M19 4v16" /><path d="M22 6 2 18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
