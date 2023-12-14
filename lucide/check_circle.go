package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CheckCircle(children ...g.Node) g.Node {
	svg := `<path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" /><path d="m9 11 3 3L22 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
