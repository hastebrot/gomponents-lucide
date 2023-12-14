package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Mic2(children ...g.Node) g.Node {
	svg := `<path d="m12 8-9.04 9.06a2.82 2.82 0 1 0 3.98 3.98L16 12" /><circle cx="17" cy="7" r="5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
