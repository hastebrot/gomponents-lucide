package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Barcode(children ...g.Node) g.Node {
	svg := `<path d="M3 5v14" /><path d="M8 5v14" /><path d="M12 5v14" /><path d="M17 5v14" /><path d="M21 5v14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
