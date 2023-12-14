package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Ligature(children ...g.Node) g.Node {
	svg := `<path d="M8 20V8c0-2.2 1.8-4 4-4 1.5 0 2.8.8 3.5 2" /><path d="M6 12h4" /><path d="M14 12h2v8" /><path d="M6 20h4" /><path d="M14 20h4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
