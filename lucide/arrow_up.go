package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUp(children ...g.Node) g.Node {
	svg := `<path d="m5 12 7-7 7 7" /><path d="M12 19V5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
