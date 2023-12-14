package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PowerCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M12 12V6" /><path d="M8 7.5A6.1 6.1 0 0 0 12 18a6 6 0 0 0 4-10.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
