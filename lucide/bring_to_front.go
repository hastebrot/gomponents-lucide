package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BringToFront(children ...g.Node) g.Node {
	svg := `<rect x="8" y="8" width="8" height="8" rx="2" /><path d="M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2" /><path d="M14 20a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
