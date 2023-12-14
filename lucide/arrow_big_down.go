package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigDown(children ...g.Node) g.Node {
	svg := `<path d="M15 6v6h4l-7 7-7-7h4V6h6z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
