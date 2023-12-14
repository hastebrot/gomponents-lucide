package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Music2(children ...g.Node) g.Node {
	svg := `<circle cx="8" cy="18" r="4" /><path d="M12 18V2l7 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
