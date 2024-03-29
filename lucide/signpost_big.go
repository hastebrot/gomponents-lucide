package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SignpostBig(children ...g.Node) g.Node {
	svg := `<path d="M10 9H4L2 7l2-2h6" /><path d="M14 5h6l2 2-2 2h-6" /><path d="M10 22V4a2 2 0 1 1 4 0v18" /><path d="M8 22h8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
