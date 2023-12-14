package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GalleryHorizontal(children ...g.Node) g.Node {
	svg := `<path d="M2 3v18" /><rect width="12" height="18" x="6" y="3" rx="2" /><path d="M22 3v18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
