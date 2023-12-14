package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Unlink2(children ...g.Node) g.Node {
	svg := `<path d="M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
