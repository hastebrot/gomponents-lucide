package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TentTree(children ...g.Node) g.Node {
	svg := `<circle cx="4" cy="4" r="2" /><path d="m14 5 3-3 3 3" /><path d="m14 10 3-3 3 3" /><path d="M17 14V2" /><path d="M17 14H7l-5 8h20Z" /><path d="M8 14v8" /><path d="m9 14 5 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
