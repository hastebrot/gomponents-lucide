package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func School(children ...g.Node) g.Node {
	svg := `<path d="m4 6 8-4 8 4" /><path d="m18 10 4 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8l4-2" /><path d="M14 22v-4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v4" /><path d="M18 5v17" /><path d="M6 5v17" /><circle cx="12" cy="9" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
