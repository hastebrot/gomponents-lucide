package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Disc2(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><circle cx="12" cy="12" r="4" /><path d="M12 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
