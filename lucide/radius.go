package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Radius(children ...g.Node) g.Node {
	svg := `<path d="M20.34 17.52a10 10 0 1 0-2.82 2.82" /><circle cx="19" cy="19" r="2" /><path d="m13.41 13.41 4.18 4.18" /><circle cx="12" cy="12" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
