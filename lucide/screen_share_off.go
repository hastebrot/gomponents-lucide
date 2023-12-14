package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ScreenShareOff(children ...g.Node) g.Node {
	svg := `<path d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3" /><path d="M8 21h8" /><path d="M12 17v4" /><path d="m22 3-5 5" /><path d="m17 3 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}