package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Kanban(children ...g.Node) g.Node {
	svg := `<path d="M6 5v11" /><path d="M12 5v6" /><path d="M18 5v14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
