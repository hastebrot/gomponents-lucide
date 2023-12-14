package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Cookie(children ...g.Node) g.Node {
	svg := `<path d="M12 2a10 10 0 1 0 10 10 4 4 0 0 1-5-5 4 4 0 0 1-5-5" /><path d="M8.5 8.5v.01" /><path d="M16 15.5v.01" /><path d="M12 12v.01" /><path d="M11 17v.01" /><path d="M7 14v.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
