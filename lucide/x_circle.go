package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func XCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m15 9-6 6" /><path d="m9 9 6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
