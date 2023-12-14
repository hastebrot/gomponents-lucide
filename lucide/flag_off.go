package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FlagOff(children ...g.Node) g.Node {
	svg := `<path d="M8 2c3 0 5 2 8 2s4-1 4-1v11" /><path d="M4 22V4" /><path d="M4 15s1-1 4-1 5 2 8 2" /><line x1="2" x2="22" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
