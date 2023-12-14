package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ListVideo(children ...g.Node) g.Node {
	svg := `<path d="M12 12H3" /><path d="M16 6H3" /><path d="M12 18H3" /><path d="m16 12 5 3-5 3v-6Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
