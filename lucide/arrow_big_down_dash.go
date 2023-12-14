package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigDownDash(children ...g.Node) g.Node {
	svg := `<path d="M15 5H9" /><path d="M15 9v3h4l-7 7-7-7h4V9h6z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
