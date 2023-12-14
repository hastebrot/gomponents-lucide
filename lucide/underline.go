package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Underline(children ...g.Node) g.Node {
	svg := `<path d="M6 4v6a6 6 0 0 0 12 0V4" /><line x1="4" x2="20" y1="20" y2="20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
