package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Expand(children ...g.Node) g.Node {
	svg := `<path d="m21 21-6-6m6 6v-4.8m0 4.8h-4.8" /><path d="M3 16.2V21m0 0h4.8M3 21l6-6" /><path d="M21 7.8V3m0 0h-4.8M21 3l-6 6" /><path d="M3 7.8V3m0 0h4.8M3 3l6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
