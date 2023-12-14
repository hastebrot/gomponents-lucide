package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cigarette(children ...g.Node) g.Node {
	svg := `<path d="M18 12H2v4h16" /><path d="M22 12v4" /><path d="M7 12v4" /><path d="M18 8c0-2.5-2-2.5-2-5" /><path d="M22 8c0-2.5-2-2.5-2-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
