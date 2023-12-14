package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Music(children ...g.Node) g.Node {
	svg := `<path d="M9 18V5l12-2v13" /><circle cx="6" cy="18" r="3" /><circle cx="18" cy="16" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
