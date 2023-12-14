package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RemoveFormatting(children ...g.Node) g.Node {
	svg := `<path d="M4 7V4h16v3" /><path d="M5 20h6" /><path d="M13 4 8 20" /><path d="m15 15 5 5" /><path d="m20 15-5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
