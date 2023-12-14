package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CheckCircle2(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m9 12 2 2 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
