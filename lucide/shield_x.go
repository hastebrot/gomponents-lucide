package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ShieldX(children ...g.Node) g.Node {
	svg := `<path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10" /><path d="m14.5 9-5 5" /><path d="m9.5 9 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
