package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ShowerHead(children ...g.Node) g.Node {
	svg := `<path d="m4 4 2.5 2.5" /><path d="M13.5 6.5a4.95 4.95 0 0 0-7 7" /><path d="M15 5 5 15" /><path d="M14 17v.01" /><path d="M10 16v.01" /><path d="M13 13v.01" /><path d="M16 10v.01" /><path d="M11 20v.01" /><path d="M17 14v.01" /><path d="M20 11v.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
