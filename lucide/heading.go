package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Heading(children ...g.Node) g.Node {
	svg := `<path d="M6 12h12" /><path d="M6 20V4" /><path d="M18 20V4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
