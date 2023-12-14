package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PlusCircle(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M8 12h8" /><path d="M12 8v8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
