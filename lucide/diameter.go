package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Diameter(children ...g.Node) g.Node {
	svg := `<circle cx="19" cy="19" r="2" /><circle cx="5" cy="5" r="2" /><path d="M6.48 3.66a10 10 0 0 1 13.86 13.86" /><path d="m6.41 6.41 11.18 11.18" /><path d="M3.66 6.48a10 10 0 0 0 13.86 13.86" />`
	return Icon(g.Group(children), g.Raw(svg))
}
