package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GitCommitHorizontal(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="3" /><line x1="3" x2="9" y1="12" y2="12" /><line x1="15" x2="21" y1="12" y2="12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
