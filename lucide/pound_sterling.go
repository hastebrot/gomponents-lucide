package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PoundSterling(children ...g.Node) g.Node {
	svg := `<path d="M18 7c0-5.333-8-5.333-8 0" /><path d="M10 7v14" /><path d="M6 21h12" /><path d="M6 13h10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
