package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Library(children ...g.Node) g.Node {
	svg := `<path d="m16 6 4 14" /><path d="M12 6v14" /><path d="M8 8v12" /><path d="M4 4v16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
