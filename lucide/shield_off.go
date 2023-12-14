package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ShieldOff(children ...g.Node) g.Node {
	svg := `<path d="M19.7 14a6.9 6.9 0 0 0 .3-2V5l-8-3-3.2 1.2" /><path d="m2 2 20 20" /><path d="M4.7 4.7 4 5v7c0 6 8 10 8 10a20.3 20.3 0 0 0 5.62-4.38" />`
	return Icon(g.Group(children), g.Raw(svg))
}
