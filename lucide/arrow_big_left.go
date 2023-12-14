package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigLeft(children ...g.Node) g.Node {
	svg := `<path d="M18 15h-6v4l-7-7 7-7v4h6v6z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
