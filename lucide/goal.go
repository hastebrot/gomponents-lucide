package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Goal(children ...g.Node) g.Node {
	svg := `<path d="M12 13V2l8 4-8 4" /><path d="M20.55 10.23A9 9 0 1 1 8 4.94" /><path d="M8 10a5 5 0 1 0 8.9 2.02" />`
	return Icon(g.Group(children), g.Raw(svg))
}
