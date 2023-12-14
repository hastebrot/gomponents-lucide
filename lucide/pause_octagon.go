package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PauseOctagon(children ...g.Node) g.Node {
	svg := `<path d="M10 15V9" /><path d="M14 15V9" /><path d="M7.714 2h8.572L22 7.714v8.572L16.286 22H7.714L2 16.286V7.714L7.714 2z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
