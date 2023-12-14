package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func DiscAlbum(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><circle cx="12" cy="12" r="5" /><path d="M12 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
