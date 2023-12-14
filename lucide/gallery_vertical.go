package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GalleryVertical(children ...g.Node) g.Node {
	svg := `<path d="M3 2h18" /><rect width="18" height="12" x="3" y="6" rx="2" /><path d="M3 22h18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
