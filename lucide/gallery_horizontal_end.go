package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GalleryHorizontalEnd(children ...g.Node) g.Node {
	svg := `<path d="M2 7v10" /><path d="M6 5v14" /><rect width="12" height="18" x="10" y="3" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
