package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileArchive(children ...g.Node) g.Node {
	svg := `<path d="M4 22V4c0-.5.2-1 .6-1.4C5 2.2 5.5 2 6 2h8.5L20 7.5V20c0 .5-.2 1-.6 1.4-.4.4-.9.6-1.4.6h-2" /><polyline points="14 2 14 8 20 8" /><circle cx="10" cy="20" r="2" /><path d="M10 7V6" /><path d="M10 12v-1" /><path d="M10 18v-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
