package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PlugZap2(children ...g.Node) g.Node {
	svg := `<path d="m13 2-2 2.5h3L12 7" /><path d="M10 14v-3" /><path d="M14 14v-3" /><path d="M11 19c-1.7 0-3-1.3-3-3v-2h8v2c0 1.7-1.3 3-3 3Z" /><path d="M12 22v-3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
