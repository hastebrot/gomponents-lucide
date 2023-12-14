package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GalleryThumbnails(children ...g.Node) g.Node {
	svg := `<rect width="18" height="14" x="3" y="3" rx="2" /><path d="M4 21h1" /><path d="M9 21h1" /><path d="M14 21h1" /><path d="M19 21h1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
