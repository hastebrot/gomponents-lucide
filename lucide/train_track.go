package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TrainTrack(children ...g.Node) g.Node {
	svg := `<path d="M2 17 17 2" /><path d="m2 14 8 8" /><path d="m5 11 8 8" /><path d="m8 8 8 8" /><path d="m11 5 8 8" /><path d="m14 2 8 8" /><path d="M7 22 22 7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
