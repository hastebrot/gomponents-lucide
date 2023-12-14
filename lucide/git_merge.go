package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GitMerge(children ...g.Node) g.Node {
	svg := `<circle cx="18" cy="18" r="3" /><circle cx="6" cy="6" r="3" /><path d="M6 21V9a9 9 0 0 0 9 9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
