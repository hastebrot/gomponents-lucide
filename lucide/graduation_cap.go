package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GraduationCap(children ...g.Node) g.Node {
	svg := `<path d="M22 10v6M2 10l10-5 10 5-10 5z" /><path d="M6 12v5c3 3 9 3 12 0v-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
