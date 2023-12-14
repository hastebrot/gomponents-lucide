package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileLock2(children ...g.Node) g.Node {
	svg := `<path d="M4 5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4" /><polyline points="14 2 14 8 20 8" /><rect width="8" height="5" x="2" y="13" rx="1" /><path d="M8 13v-2a2 2 0 1 0-4 0v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
