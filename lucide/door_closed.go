package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func DoorClosed(children ...g.Node) g.Node {
	svg := `<path d="M18 20V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14" /><path d="M2 20h20" /><path d="M14 12v.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
