package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Euro(children ...g.Node) g.Node {
	svg := `<path d="M4 10h12" /><path d="M4 14h9" /><path d="M19 6a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8 2 0 3.8-.8 5.2-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
