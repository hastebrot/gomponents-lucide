package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GitGraph(children ...g.Node) g.Node {
	svg := `<circle cx="5" cy="6" r="3" /><path d="M5 9v6" /><circle cx="5" cy="18" r="3" /><path d="M12 3v18" /><circle cx="19" cy="6" r="3" /><path d="M16 15.7A9 9 0 0 0 19 9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
