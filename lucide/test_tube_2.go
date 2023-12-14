package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TestTube2(children ...g.Node) g.Node {
	svg := `<path d="M21 7 6.82 21.18a2.83 2.83 0 0 1-3.99-.01v0a2.83 2.83 0 0 1 0-4L17 3" /><path d="m16 2 6 6" /><path d="M12 16H4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
