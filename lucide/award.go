package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Award(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="8" r="6" /><path d="M15.477 12.89 17 22l-5-3-5 3 1.523-9.11" />`
	return Icon(g.Group(children), g.Raw(svg))
}
