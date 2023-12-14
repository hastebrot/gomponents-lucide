package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func LampCeiling(children ...g.Node) g.Node {
	svg := `<path d="M12 2v5" /><path d="M6 7h12l4 9H2l4-9Z" /><path d="M9.17 16a3 3 0 1 0 5.66 0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
