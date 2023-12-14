package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Shrink(children ...g.Node) g.Node {
	svg := `<path d="m15 15 6 6m-6-6v4.8m0-4.8h4.8" /><path d="M9 19.8V15m0 0H4.2M9 15l-6 6" /><path d="M15 4.2V9m0 0h4.8M15 9l6-6" /><path d="M9 4.2V9m0 0H4.2M9 9 3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
