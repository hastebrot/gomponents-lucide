package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowBigLeftDash(children ...g.Node) g.Node {
	svg := `<path d="M19 15V9" /><path d="M15 15h-3v4l-7-7 7-7v4h3v6z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
