package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PauseCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><line x1="10" x2="10" y1="15" y2="9" /><line x1="14" x2="14" y1="15" y2="9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
