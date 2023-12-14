package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Loader2(children ...g.Node) g.Node {
	svg := `<path d="M21 12a9 9 0 1 1-6.219-8.56" />`
	return Icon(g.Group(children), g.Raw(svg))
}
