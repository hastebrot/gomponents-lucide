package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BellPlus(children ...g.Node) g.Node {
	svg := `<path d="M19.3 14.8C20.1 16.4 21 17 21 17H3s3-2 3-9c0-3.3 2.7-6 6-6 1 0 1.9.2 2.8.7" /><path d="M10.3 21a1.94 1.94 0 0 0 3.4 0" /><path d="M15 8h6" /><path d="M18 5v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
