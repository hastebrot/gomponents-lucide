package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowDown(children ...g.Node) g.Node {
	svg := `<path d="M12 5v14" /><path d="m19 12-7 7-7-7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
