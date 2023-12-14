package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GalleryVerticalEnd(children ...g.Node) g.Node {
	svg := `<path d="M7 2h10" /><path d="M5 6h14" /><rect width="18" height="12" x="3" y="10" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
