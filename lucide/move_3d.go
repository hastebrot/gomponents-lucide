package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Move3d(children ...g.Node) g.Node {
	svg := `<path d="M5 3v16h16" /><path d="m5 19 6-6" /><path d="m2 6 3-3 3 3" /><path d="m18 16 3 3-3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
