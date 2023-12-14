package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Bold(children ...g.Node) g.Node {
	svg := `<path d="M14 12a4 4 0 0 0 0-8H6v8" /><path d="M15 20a4 4 0 0 0 0-8H6v8Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
