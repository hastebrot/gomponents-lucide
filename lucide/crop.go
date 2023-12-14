package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Crop(children ...g.Node) g.Node {
	svg := `<path d="M6 2v14a2 2 0 0 0 2 2h14" /><path d="M18 22V8a2 2 0 0 0-2-2H2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
