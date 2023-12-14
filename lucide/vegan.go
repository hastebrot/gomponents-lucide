package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Vegan(children ...g.Node) g.Node {
	svg := `<path d="M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14" /><path d="M16 8c4 0 6-2 6-6-4 0-6 2-6 6" /><path d="M17.41 3.6a10 10 0 1 0 3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
