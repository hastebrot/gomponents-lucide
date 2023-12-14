package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tornado(children ...g.Node) g.Node {
	svg := `<path d="M21 4H3" /><path d="M18 8H6" /><path d="M19 12H9" /><path d="M16 16h-6" /><path d="M11 20H9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
