package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StopCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><rect width="6" height="6" x="9" y="9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
