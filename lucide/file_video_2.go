package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileVideo2(children ...g.Node) g.Node {
	svg := `<path d="M4 8V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4" /><polyline points="14 2 14 8 20 8" /><path d="m10 15.5 4 2.5v-6l-4 2.5" /><rect width="8" height="6" x="2" y="12" rx="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
